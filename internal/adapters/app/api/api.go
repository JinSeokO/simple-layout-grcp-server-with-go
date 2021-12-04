package api

import (
	"youtube/internal/ports"
)

type Adapter struct {
	db    ports.DBPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DBPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db, arith}
}

func (adapter Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := adapter.arith.Addition(a, b)
	if err != nil {
		return 0, nil
	}

	err = adapter.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (adapter Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := adapter.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.db.AddToHistory(answer, "Subtraction")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (adapter Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := adapter.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = adapter.db.AddToHistory(answer, "Multiple")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (adapter Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := adapter.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.db.AddToHistory(answer, "Division")
	if err != nil {
		return 0, err
	}
	return answer, nil
}
