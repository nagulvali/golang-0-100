package app

import (
	"go.temporal.io/sdk/workflow"
)

type (
    CartItem struct {
        ProductId int
        Quantity  int
    }

    CartState struct {
        Items []CartItem
        Email string
    }
)

func CartWorkflow(ctx workflow.Context, state CartState) error {
    logger := workflow.GetLogger(ctx)


		channel := workflow.GetSignalChannel(ctx, "cartMessages")
		selector := workflow.NewSelector(ctx)

		selector.AddReceive(channel, func(c workflow.ReceiveChannel, _ bool) {
				var signal interface{}
				c.Receive(ctx, &signal)

				var routeSignal RouteSignal
				err := mapstructure.Decode(signal, &routeSignal)
				if err != nil {
						logger.Error("Invalid signal type %v", err)
						return
				}

				switch {
				case routeSignal.Route == RouteTypes.ADD_TO_CART:
						var message AddToCartSignal
						err := mapstructure.Decode(signal, &message)
						if err != nil {
								logger.Error("Invalid signal type %v", err)
								return
						}
						AddToCart(&state, message.Item)

				case routeSignal.Route == RouteTypes.REMOVE_FROM_CART:
						var message RemoveFromCartSignal
						err := mapstructure.Decode(signal, &message)
						if err != nil {
								logger.Error("Invalid signal type %v", err)
								return
						}
						RemoveFromCart(&state, message.Item)
		})

		// Stop blocking once one condition is satisfied
		for {
				selector.Select(ctx)
		}

		
		
}