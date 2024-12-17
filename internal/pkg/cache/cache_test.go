package cache

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	t.Parallel()

	t.Run("Value not exists", func(t *testing.T) {
		value, exists := Get("test1")
		require.Nil(t, value)
		require.False(t, exists)
	})

	t.Run("Value exists", func(t *testing.T) {
		key := gofakeit.Word()
		value := gofakeit.Word()
		Set(key, value, 10*time.Second)
		iValue, exists := Get(key)
		require.NotNil(t, iValue)
		require.True(t, exists)
		sValue, ok := iValue.(string)
		require.True(t, ok)
		require.Equal(t, sValue, value)
	})

	t.Run("Value is expired", func(t *testing.T) {
		key := gofakeit.Word()
		value := gofakeit.Word()
		Set(key, value, time.Millisecond)
		time.Sleep(10 * time.Millisecond)
		iValue, exists := Get(key)
		require.Nil(t, iValue)
		require.False(t, exists)
	})

	t.Run("After delete", func(t *testing.T) {
		key := gofakeit.Word()
		value := gofakeit.Word()
		Set(key, value, 10*time.Second)
		Delete(key)
		iValue, exists := Get(key)
		require.Nil(t, iValue)
		require.False(t, exists)
	})
}
