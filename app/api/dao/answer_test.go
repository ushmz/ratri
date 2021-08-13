package dao

import (
	"fmt"
	"testing"

	"baltard/api/models"

	"github.com/stretchr/testify/assert"
)

func TestSubmitTaskAnswer(t *testing.T) {
	tests := []struct {
		ans      models.Answer
		expected error
	}{
		{
			ans: models.Answer{
				UserId:      44422,
				TaskId:      5,
				ConditionId: 5,
				Answer:      "otameshi answer",
				Reason:      "otameshi reason",
			},
			expected: nil},
		{
			ans: models.Answer{
				UserId:      4422,
				TaskId:      5,
				ConditionId: 5,
				Answer:      "お試し アンサー",
				Reason:      "お試し リーズン",
			},
			expected: nil},
		{
			ans: models.Answer{
				UserId:      4423,
				TaskId:      6,
				ConditionId: 5,
				Answer:      "akaaaan answer",
				Reason:      "hoge); SELECT * FROM users; -- ",
			},
			expected: nil},
	}

	for idx, test := range tests {
		err := answerDao.SubmitTaskAnswer(&test.ans)

		assert.Equal(
			t,
			test.expected,
			err,
			fmt.Sprintf("Testcase index %d", idx))
	}
}