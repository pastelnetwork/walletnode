package ticket

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/labstack/echo/v4"
	"github.com/pastelnetwork/go-commons/log"
	"github.com/pastelnetwork/go-pastel"
	"github.com/pastelnetwork/walletnode/internal/restserver"
	"golang.org/x/sync/errgroup"
)

type TicketProc struct {
	pslNode *pastel.Client
}

type WSMessage struct {
	op  ws.OpCode
	msg []byte
}

func NewTicketProc(psl *pastel.Client) *TicketProc {
	return &TicketProc{
		pslNode: psl,
	}
}

func (p *TicketProc) RegisterArtTicket(c echo.Context) error {
	conn, _, _, err := ws.UpgradeHTTP(c.Request(), c.Response().Writer)
	if err != nil {
		return err
	}

	cc := c.(*restserver.RESTServerContext)
	//cc.Jobs.Go( func() error {
	go func() error {
		defer conn.Close()

		log.Infoln("New Ticket Processor started")

		eg, ctx := errgroup.WithContext(cc.AppCtx)

		messages := make(chan WSMessage)

		eg.Go(func() error {
			log.Infoln("NTP WS Listener started")
			for {
				log.Infoln("Waiting for message")
				msg, op, err := wsutil.ReadClientData(conn)
				log.Infoln("NTP WS Worker exiting - error and signal check")
				if err != nil {
					log.Errorf("Error in New Ticket Processor Listener - %w", err)
					return err
				}
				select {
				case <-ctx.Done():
					log.Infoln("NTP WS Listener exiting")
					return nil
				case messages <- WSMessage{op, msg}:
					continue
				}
			}
		})
		eg.Go(func() error {
			log.Infoln("NTP WS Worker started")
			for {
				select {
				case <-ctx.Done():
					log.Infoln("NTP WS Worker exiting")
					err = wsutil.WriteServerMessage(conn, ws.OpClose, nil)
					return nil
				case msg := <-messages:
					log.Infof("Got message - %s (opcode - %c)", msg.msg, msg.op)
					err = wsutil.WriteServerMessage(conn, msg.op, msg.msg)
					if err != nil {
						log.Errorf("Error in New Ticket Processor Listener - %s", err)
						return err
					}
				}
			}
		})
		if err := eg.Wait(); err != nil {
			log.Errorf("Error in New Ticket Processor - %s", err)
			return nil
		}

		log.Infoln("New Ticket Processor exiting")
		return nil
	}()
	return nil
}
