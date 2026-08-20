package charging

import "errors"

var ErrInvalidStation = errors.New("invalid station")

type StationBatch struct {
	stored []string
	active int
}
type stationTx struct {
	owner  *StationBatch
	staged []string
	closed bool
}

func (s *StationBatch) begin() *stationTx { s.active++; return &stationTx{owner: s} }
func (tx *stationTx) Close() {
	if !tx.closed {
		tx.closed = true
		tx.owner.active--
	}
}
func (tx *stationTx) Rollback() { tx.staged = nil; tx.Close() }
func (tx *stationTx) Commit()   { tx.owner.stored = append(tx.owner.stored, tx.staged...); tx.Close() }
func (s *StationBatch) Register(names []string) error {
	tx := s.begin()
	defer tx.Rollback()
	for _, name := range names {
		if name == "invalid" {
			return ErrInvalidStation
		}
		tx.staged = append(tx.staged, name)
	}
	tx.Commit()
	return nil
}
func (s *StationBatch) Stats() (int, int) { return len(s.stored), s.active }
