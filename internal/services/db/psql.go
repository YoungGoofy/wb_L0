package db

import (
	"context"
	"github.com/YoungGoofy/WB_L0/internal/models"
	"github.com/YoungGoofy/WB_L0/internal/services/postgresql"
)

type OrderRepository struct {
	client postgresql.Client
}

func NewOrderRepository(client postgresql.Client) *OrderRepository {
	return &OrderRepository{client: client}
}

func (r *OrderRepository) GetIds(ctx context.Context) ([]string, error) {
	var uids []string
	rows, err := r.client.Query(ctx, getUID)
	if err != nil {
		return nil, err
	}
	var uid models.Orders
	for rows.Next() {
		err = rows.Scan(&uid.OrderUID)
		if err != nil {
			return nil, err
		}
		uids = append(uids, uid.OrderUID)
	}
	return uids, nil
}

func (r *OrderRepository) CreateOrder(ctx context.Context, orders *models.Orders) (*models.Orders, error) {
	var newOrder models.Orders

	if err := r.client.QueryRow(ctx, queryCreateOrder,
		orders.OrderUID,
		orders.TrackNumber,
		orders.Entry,
		orders.Locale,
		orders.InternalSignature,
		orders.CustomerId,
		orders.DeliveryService,
		orders.Shardkey,
		orders.SmId,
		//orders.DateCreated,
		orders.OofShard).Scan(
		&newOrder.OrderUID,
	); err != nil {
		return nil, err
	}
	return &newOrder, nil
}

func (r *OrderRepository) Create(ctx context.Context, order *models.Orders) error {

	_, err := r.client.Exec(ctx, queryCreateOrder,
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		order.Locale,
		order.InternalSignature,
		order.CustomerId,
		order.DeliveryService,
		order.Shardkey,
		order.SmId,
		//order.DateCreated,
		order.OofShard)
	if err != nil {
		return err
	}

	_, err = r.client.Exec(ctx, queryCreateDelivery,
		order.OrderUID,
		order.Delivery.Name,
		order.Delivery.Phone,
		order.Delivery.Zip,
		order.Delivery.City,
		order.Delivery.Address,
		order.Delivery.Region,
		order.Delivery.Email)
	if err != nil {
		return err
	}

	_, err = r.client.Exec(ctx, queryCreatePayment,
		order.OrderUID,
		order.Payment.Transaction,
		order.Payment.RequestId,
		order.Payment.Currency,
		order.Payment.Provider,
		order.Payment.Amount,
		order.Payment.PaymentDt,
		order.Payment.Bank,
		order.Payment.DeliveryCost,
		order.Payment.GoodsTotal,
		order.Payment.CustomFee)
	if err != nil {
		return err
	}

	for _, v := range order.Items {
		_, err = r.client.Exec(ctx, queryCreateItem,
			order.OrderUID,
			v.ChrtId,
			v.TrackNumber,
			v.Price,
			v.Rid,
			v.Name,
			v.Sale,
			v.Size,
			v.TotalPrice,
			v.NmId,
			v.Brand,
			v.Status)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *OrderRepository) GetOrderById(ctx context.Context, uid string) (models.Orders, error) {
	var order models.Orders
	if err := r.client.QueryRow(ctx, getOrderByOrderUIDQuery, uid).Scan(
		&order.OrderUID,
		&order.TrackNumber,
		&order.Entry,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerId,
		&order.DeliveryService,
		&order.Shardkey,
		&order.SmId,
		&order.DateCreated,
		&order.OofShard); err != nil {
		return order, err
	}
	return order, nil
}

func (r *OrderRepository) GetFullById(ctx context.Context, uid string) (*models.Orders, error) {
	var order models.Orders

	err := r.client.QueryRow(ctx, getOrderByOrderUIDQuery, uid).Scan(
		&order.OrderUID,
		&order.TrackNumber,
		&order.Entry,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerId,
		&order.DeliveryService,
		&order.Shardkey,
		&order.SmId,
		&order.DateCreated,
		&order.OofShard)
	if err != nil {
		return nil, err
	}

	err = r.client.QueryRow(ctx, getDeliveryByOrderUIDQuery, uid).Scan(
		&order.Delivery.Name,
		&order.Delivery.Phone,
		&order.Delivery.Zip,
		&order.Delivery.City,
		&order.Delivery.Address,
		&order.Delivery.Region,
		&order.Delivery.Email)
	if err != nil {
		return nil, err
	}

	err = r.client.QueryRow(ctx, getPaymentByOrderUIDQuery, uid).Scan(
		&order.Payment.Transaction,
		&order.Payment.RequestId,
		&order.Payment.Currency,
		&order.Payment.Provider,
		&order.Payment.Amount,
		&order.Payment.PaymentDt,
		&order.Payment.Bank,
		&order.Payment.DeliveryCost,
		&order.Payment.GoodsTotal,
		&order.Payment.CustomFee)
	if err != nil {
		return nil, err
	}

	rows, err := r.client.Query(ctx, getItemsByOrderUIDQuery, uid)
	if err != nil {
		return nil, err
	}
	item := models.Item{}

	for rows.Next() {
		err = rows.Scan(
			&item.ChrtId,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmId,
			&item.Brand,
			&item.Status)
		if err != nil {
			return nil, err
		}
		order.Items = append(order.Items, item)
	}
	return &order, nil
}
