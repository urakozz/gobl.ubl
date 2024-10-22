package utog

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseUtoGOrdering(t *testing.T) {
	t.Run("UBL_example1.xml", func(t *testing.T) {
		doc, err := LoadTestXMLDoc("UBL_example1.xml")
		require.NoError(t, err)

		conversor := NewConversor()
		inv, err := conversor.NewInvoice(doc)
		require.NoError(t, err)
		ordering := inv.Ordering
		assert.NotNil(t, ordering)
	})

}