package cabify_challenge

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestDefaultPrice(t *testing.T) {
  p := DefaultPrice { pricePerUnit: 1.95 }
  assert.Equal(t, 0.0, p.Calculate(0))
  assert.Equal(t, 1.95, p.Calculate(1))
  assert.Equal(t, 3.90, p.Calculate(2))
}

func TestGetTwoPayOnePrice(t *testing.T) {
  p := GetTwoPayOnePrice { pricePerUnit: 1.95 }
  assert.Equal(t, 0.0, p.Calculate(0))
  assert.Equal(t, 1.95, p.Calculate(1))
  assert.Equal(t, 1.95, p.Calculate(2))
  assert.Equal(t, 3.90, p.Calculate(3))
  assert.Equal(t, 3.90, p.Calculate(4))
}

func TestBulkDiscountPrice(t *testing.T) {
  p := BulkDiscountPrice { pricePerUnit: 1.95, discountPricePerUnit: 1.75, bulkMinUnits: 3 }
  assert.Equal(t, 0.0, p.Calculate(0))
  assert.Equal(t, 1.95, p.Calculate(1))
  assert.Equal(t, 3.90, p.Calculate(2))
  assert.Equal(t, 5.25, p.Calculate(3))
  assert.Equal(t, 7.00, p.Calculate(4))
}