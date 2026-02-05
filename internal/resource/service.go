package resource

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/bencoronard/demo-go-common-libs/constant"
	"github.com/bencoronard/demo-go-common-libs/dto"
	"github.com/bencoronard/demo-go-common-libs/orm"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type resourceService interface {
	listResources(ctx context.Context, page dto.Pageable, claims jwt.MapClaims) (dto.Slice[resource], error)
	retrieveResource(ctx context.Context, id uint, claims jwt.MapClaims) (resource, error)
	createResource(ctx context.Context, dto resource, claims jwt.MapClaims) (uint, error)
	updateResource(ctx context.Context, id uint, dto resource, claims jwt.MapClaims) error
	deleteResource(ctx context.Context, id uint, claims jwt.MapClaims) error
}

type resourceServiceImpl struct {
	t orm.TransactionManager
	r resourceRepo
}

func NewResourceService(t orm.TransactionManager, r resourceRepo) resourceService {
	return &resourceServiceImpl{
		t: t,
		r: r,
	}
}

func (s *resourceServiceImpl) listResources(ctx context.Context, page dto.Pageable, claims jwt.MapClaims) (dto.Slice[resource], error) {
	val, ok := claims["list_resources"]
	if !ok || val == nil {
		return dto.Slice[resource]{}, fmt.Errorf("%w: not allowed to list resources", constant.ErrInsufficientPermission)
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return dto.Slice[resource]{}, err
	}

	createdBy, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return dto.Slice[resource]{}, err
	}

	var slice dto.Slice[resource]

	err = s.t.Transactional(ctx, func(tx *gorm.DB) error {
		slice, err = s.r.findAllByCreatedBy(tx, page, uint(createdBy))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return dto.Slice[resource]{}, err
	}

	return slice, nil
}

func (s *resourceServiceImpl) retrieveResource(ctx context.Context, id uint, claims jwt.MapClaims) (resource, error) {
	val, ok := claims["read_resource"]
	if !ok || val == nil {
		return resource{}, fmt.Errorf("%w: not allowed to read resource", constant.ErrInsufficientPermission)
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return resource{}, err
	}

	createdBy, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return resource{}, err
	}

	var res resource

	err = s.t.Transactional(ctx, func(tx *gorm.DB) error {
		res, err = s.r.findByIdAndCreatedBy(tx, id, uint(createdBy))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("%w: resource id: %d not found", ErrResourceNotFound, id)
			}
			return err
		}
		return nil
	})
	if err != nil {
		return resource{}, err
	}

	return res, err
}

func (s *resourceServiceImpl) createResource(ctx context.Context, dto resource, claims jwt.MapClaims) (uint, error) {
	val, ok := claims["create_resource"]
	if !ok || val == nil {
		return 0, fmt.Errorf("%w: not allowed to create resource", constant.ErrInsufficientPermission)
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return 0, err
	}

	createdBy, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return 0, err
	}

	var res resource
	res.CreatedBy = uint(createdBy)
	res.TextField = dto.TextField
	res.NumberField = dto.NumberField
	res.BooleanField = dto.BooleanField

	err = s.t.Transactional(ctx, func(tx *gorm.DB) error {
		if err := s.r.save(tx, &res); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return res.ID, err
}

func (s *resourceServiceImpl) updateResource(ctx context.Context, id uint, dto resource, claims jwt.MapClaims) error {
	val, ok := claims["update_resource"]
	if !ok || val == nil {
		return fmt.Errorf("%w: not allowed to update resource", constant.ErrInsufficientPermission)
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return err
	}

	createdBy, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return err
	}

	err = s.t.Transactional(ctx, func(tx *gorm.DB) error {
		resource, err := s.r.findByIdAndCreatedBy(tx, id, uint(createdBy))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("%w: resource id: %d not found", ErrResourceNotFound, id)
			}
			return err
		}

		resource.TextField = dto.TextField
		resource.NumberField = dto.NumberField
		resource.BooleanField = dto.BooleanField

		if err := s.r.save(tx, &resource); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *resourceServiceImpl) deleteResource(ctx context.Context, id uint, claims jwt.MapClaims) error {
	val, ok := claims["delete_resource"]
	if !ok || val == nil {
		return fmt.Errorf("%w: not allowed to delete resource", constant.ErrInsufficientPermission)
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return err
	}

	createdBy, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return err
	}

	err = s.t.Transactional(ctx, func(tx *gorm.DB) error {
		resource, err := s.r.findByIdAndCreatedBy(tx, id, uint(createdBy))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("%w: resource id: %d not found", ErrResourceNotFound, id)
			}
			return err
		}

		if err := s.r.delete(tx, resource); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
