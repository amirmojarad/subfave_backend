// Code generated by entc, DO NOT EDIT.

package wordnode

import (
	"interface_project/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
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
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
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
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// IsPreposition applies equality check predicate on the "is_preposition" field. It's identical to IsPrepositionEQ.
func IsPreposition(v bool) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsPreposition), v))
	})
}

// Occurence applies equality check predicate on the "occurence" field. It's identical to OccurenceEQ.
func Occurence(v int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOccurence), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.WordNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WordNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.WordNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WordNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// IsPrepositionEQ applies the EQ predicate on the "is_preposition" field.
func IsPrepositionEQ(v bool) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsPreposition), v))
	})
}

// IsPrepositionNEQ applies the NEQ predicate on the "is_preposition" field.
func IsPrepositionNEQ(v bool) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsPreposition), v))
	})
}

// OccurenceEQ applies the EQ predicate on the "occurence" field.
func OccurenceEQ(v int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOccurence), v))
	})
}

// OccurenceNEQ applies the NEQ predicate on the "occurence" field.
func OccurenceNEQ(v int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOccurence), v))
	})
}

// OccurenceIn applies the In predicate on the "occurence" field.
func OccurenceIn(vs ...int) predicate.WordNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WordNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOccurence), v...))
	})
}

// OccurenceNotIn applies the NotIn predicate on the "occurence" field.
func OccurenceNotIn(vs ...int) predicate.WordNode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WordNode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOccurence), v...))
	})
}

// OccurenceGT applies the GT predicate on the "occurence" field.
func OccurenceGT(v int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOccurence), v))
	})
}

// OccurenceGTE applies the GTE predicate on the "occurence" field.
func OccurenceGTE(v int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOccurence), v))
	})
}

// OccurenceLT applies the LT predicate on the "occurence" field.
func OccurenceLT(v int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOccurence), v))
	})
}

// OccurenceLTE applies the LTE predicate on the "occurence" field.
func OccurenceLTE(v int) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOccurence), v))
	})
}

// OccurenceIsNil applies the IsNil predicate on the "occurence" field.
func OccurenceIsNil() predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOccurence)))
	})
}

// OccurenceNotNil applies the NotNil predicate on the "occurence" field.
func OccurenceNotNil() predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOccurence)))
	})
}

// HasWords applies the HasEdge predicate on the "words" edge.
func HasWords() predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WordsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WordsTable, WordsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWordsWith applies the HasEdge predicate on the "words" edge with a given conditions (other predicates).
func HasWordsWith(preds ...predicate.Word) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WordsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WordsTable, WordsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFile applies the HasEdge predicate on the "file" edge.
func HasFile() predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FileTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FileTable, FileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFileWith applies the HasEdge predicate on the "file" edge with a given conditions (other predicates).
func HasFileWith(preds ...predicate.FileEntity) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FileInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FileTable, FileColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WordNode) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WordNode) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WordNode) predicate.WordNode {
	return predicate.WordNode(func(s *sql.Selector) {
		p(s.Not())
	})
}
