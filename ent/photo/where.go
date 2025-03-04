// Code generated by entc, DO NOT EDIT.

package photo

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pavelz/insta-go/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.In(s.C(FieldID), v...))
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.NotIn(s.C(FieldID), v...))
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldID), id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldID), id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldID), id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldID), id))
		},
	)
}

// Lat applies equality check predicate on the "lat" field. It's identical to LatEQ.
func Lat(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldLat), v))
		},
	)
}

// Lng applies equality check predicate on the "lng" field. It's identical to LngEQ.
func Lng(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldLng), v))
		},
	)
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v []byte) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldImage), v))
		},
	)
}

// Fielname applies equality check predicate on the "fielname" field. It's identical to FielnameEQ.
func Fielname(v string) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldFielname), v))
		},
	)
}

// LatEQ applies the EQ predicate on the "lat" field.
func LatEQ(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldLat), v))
		},
	)
}

// LatNEQ applies the NEQ predicate on the "lat" field.
func LatNEQ(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldLat), v))
		},
	)
}

// LatIn applies the In predicate on the "lat" field.
func LatIn(vs ...float64) predicate.Photo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Photo(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldLat), v...))
		},
	)
}

// LatNotIn applies the NotIn predicate on the "lat" field.
func LatNotIn(vs ...float64) predicate.Photo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Photo(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldLat), v...))
		},
	)
}

// LatGT applies the GT predicate on the "lat" field.
func LatGT(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldLat), v))
		},
	)
}

// LatGTE applies the GTE predicate on the "lat" field.
func LatGTE(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldLat), v))
		},
	)
}

// LatLT applies the LT predicate on the "lat" field.
func LatLT(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldLat), v))
		},
	)
}

// LatLTE applies the LTE predicate on the "lat" field.
func LatLTE(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldLat), v))
		},
	)
}

// LngEQ applies the EQ predicate on the "lng" field.
func LngEQ(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldLng), v))
		},
	)
}

// LngNEQ applies the NEQ predicate on the "lng" field.
func LngNEQ(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldLng), v))
		},
	)
}

// LngIn applies the In predicate on the "lng" field.
func LngIn(vs ...float64) predicate.Photo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Photo(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldLng), v...))
		},
	)
}

// LngNotIn applies the NotIn predicate on the "lng" field.
func LngNotIn(vs ...float64) predicate.Photo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Photo(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldLng), v...))
		},
	)
}

// LngGT applies the GT predicate on the "lng" field.
func LngGT(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldLng), v))
		},
	)
}

// LngGTE applies the GTE predicate on the "lng" field.
func LngGTE(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldLng), v))
		},
	)
}

// LngLT applies the LT predicate on the "lng" field.
func LngLT(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldLng), v))
		},
	)
}

// LngLTE applies the LTE predicate on the "lng" field.
func LngLTE(v float64) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldLng), v))
		},
	)
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v []byte) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldImage), v))
		},
	)
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v []byte) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldImage), v))
		},
	)
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...[]byte) predicate.Photo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Photo(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldImage), v...))
		},
	)
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...[]byte) predicate.Photo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Photo(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldImage), v...))
		},
	)
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v []byte) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldImage), v))
		},
	)
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v []byte) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldImage), v))
		},
	)
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v []byte) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldImage), v))
		},
	)
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v []byte) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldImage), v))
		},
	)
}

// ImageIsNil applies the IsNil predicate on the "image" field.
func ImageIsNil() predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldImage)))
		},
	)
}

// ImageNotNil applies the NotNil predicate on the "image" field.
func ImageNotNil() predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldImage)))
		},
	)
}

// FielnameEQ applies the EQ predicate on the "fielname" field.
func FielnameEQ(v string) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldFielname), v))
		},
	)
}

// FielnameNEQ applies the NEQ predicate on the "fielname" field.
func FielnameNEQ(v string) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldFielname), v))
		},
	)
}

// FielnameIn applies the In predicate on the "fielname" field.
func FielnameIn(vs ...string) predicate.Photo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Photo(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldFielname), v...))
		},
	)
}

// FielnameNotIn applies the NotIn predicate on the "fielname" field.
func FielnameNotIn(vs ...string) predicate.Photo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Photo(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldFielname), v...))
		},
	)
}

// FielnameGT applies the GT predicate on the "fielname" field.
func FielnameGT(v string) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldFielname), v))
		},
	)
}

// FielnameGTE applies the GTE predicate on the "fielname" field.
func FielnameGTE(v string) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldFielname), v))
		},
	)
}

// FielnameLT applies the LT predicate on the "fielname" field.
func FielnameLT(v string) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldFielname), v))
		},
	)
}

// FielnameLTE applies the LTE predicate on the "fielname" field.
func FielnameLTE(v string) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldFielname), v))
		},
	)
}

// FielnameContains applies the Contains predicate on the "fielname" field.
func FielnameContains(v string) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldFielname), v))
		},
	)
}

// FielnameHasPrefix applies the HasPrefix predicate on the "fielname" field.
func FielnameHasPrefix(v string) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldFielname), v))
		},
	)
}

// FielnameHasSuffix applies the HasSuffix predicate on the "fielname" field.
func FielnameHasSuffix(v string) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldFielname), v))
		},
	)
}

// FielnameEqualFold applies the EqualFold predicate on the "fielname" field.
func FielnameEqualFold(v string) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldFielname), v))
		},
	)
}

// FielnameContainsFold applies the ContainsFold predicate on the "fielname" field.
func FielnameContainsFold(v string) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldFielname), v))
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Photo) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Photo) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Photo) predicate.Photo {
	return predicate.Photo(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
