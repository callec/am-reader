package basicservice

import (
	"context"
	"fmt"
	"mag"
	"mag/service/db"
	"time"

	"github.com/google/uuid"
)

type basicService struct {
	r *db.Queries
}

func (b *basicService) AddMagazine(
	ctx context.Context,
	no int,
	t time.Time,
	loc string,
) error {
	param := db.AddMagazineParams{
		Number:   int64(no),
		Date:     t.Unix(),
		Location: loc,
	}
	res, err := b.r.AddMagazine(ctx, param)
	if err != nil {
		return fmt.Errorf("AddMagazine: Error reading from database: %w", err)
	}
	_, err = res.LastInsertId()
	if err != nil {
		return fmt.Errorf("AddMagazine: Error adding magazine: %w", err)
	}
	return nil
}

func (b *basicService) GetMagazine(ctx context.Context, mid uuid.UUID) (*mag.Magazine, error) {
	m, err := b.r.GetMagazine(ctx, mid.String())
	if err != nil {
		return nil, fmt.Errorf("GetMagazine: Error reading from database: %w", err)
	}

	nm, err := mkMag(&m)
	if err != nil {
		return nil, fmt.Errorf("GetMagazine: Error parsing uuid: %w", err)
	}
	return nm, nil
}

func (b *basicService) GetMagazineByNumber(ctx context.Context, mno int) (*mag.Magazine, error) {
	m, err := b.r.GetMagazineByNumber(ctx, int64(mno))
	if err != nil {
		return nil, fmt.Errorf("GetMagazine: Error reading from database: %w", err)
	}

	nm, err := mkMag(&m)
	if err != nil {
		return nil, fmt.Errorf("GetMagazine: Error parsing uuid: %w", err)
	}
	return nm, nil
}

func mkMag(m *db.Magazine) (*mag.Magazine, error) {
	mid, err := uuid.Parse(m.ID)
	if err != nil {
		return nil, err
	}
	return &mag.Magazine{
		Id:       mid,
		Date:     time.Unix(m.Date, 0),
		Number:   int(m.Number),
		Location: m.Location,
	}, nil
}

func (b *basicService) ListMagazines(
	ctx context.Context,
	limit int,
	offset int,
) ([]*mag.Magazine, error) {
	ms, err := b.r.ListMagazines(ctx, db.ListMagazinesParams{
		Limit:  int64(limit),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, fmt.Errorf("ListMagazines: Error reading from database: %w", err)
	}
	var nms []*mag.Magazine
	for _, m := range ms {
		nm, err := mkMag(&m)
		if err != nil {
			return nil, fmt.Errorf("ListMagazines: Error parsing uuid: %w", err)
		}
		nms = append(nms, nm)
	}
	return nms, nil
}

func (b *basicService) RemoveMagazine(ctx context.Context, mid uuid.UUID) error {
	return b.r.RemoveMagazine(ctx, mid.String())
}

func mkUser(u *db.User) (*mag.User, error) {
	uid, err := uuid.Parse(u.ID)
	if err != nil {
		return nil, err
	}
	return &mag.User{
		Id:           uid,
		Username:     "currently unused",
		PasswordHash: u.Pwd,
		Created:      time.Unix(u.Created, 0),
		LastOnline:   time.Unix(u.Lastonline, 0),
	}, nil
}

func (b *basicService) GetUser(ctx context.Context, uid uuid.UUID) (*mag.User, error) {
	u, err := b.r.GetUser(ctx, uid.String())
	if err != nil {
		return nil, fmt.Errorf("GetUser: Error reading from database: %w", err)
	}

	nu, err := mkUser(&u)
	if err != nil {
		return nil, fmt.Errorf("GetUser: Error parsing uuid: %w", err)
	}
	return nu, nil
}

func (b *basicService) GetUserByName(ctx context.Context, uname string) (*mag.User, error) {
	u, err := b.r.GetUid(ctx, uname)
	if err != nil {
		return nil, fmt.Errorf("GetUserByName: Error reading from database: %w", err)
	}

	uid, err := uuid.Parse(u)
	return b.GetUser(ctx, uid)
}

func (b *basicService) RegisterUser(ctx context.Context, uname string, pwd string) error {
	id, err := b.r.RegisterUser(ctx, pwd)
	if err != nil {
		return fmt.Errorf("RegisterUser: Error reading from database: %w", err)
	}

	_, err = b.r.AddUName(ctx, db.AddUNameParams{Uid: id, Uname: uname})
	if err != nil {
		return fmt.Errorf("RegisterUser: Error adding nickname: %w", err)
	}

	return nil
}

func CreateBasicService(r *db.Queries) *basicService {
	return &basicService{
		r: r,
	}
}
