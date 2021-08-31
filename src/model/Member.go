package model

type Member struct {
	memberId string "사용자 아이디"

	memberPw string "사용자 비밀번호"

	memberName string "사용자 이름"

	memberAge int "사용자 나이"

	memberGender string "사용자 성별"

	memberAddr1 string "사용자 주소1"

	memberAddr2 string "사용자 상세주소"
}

func (m *Member) SetMemberId(memberId string) {
	m.memberId = memberId
}

func (m *Member) MemberId() string {
	return m.memberId
}
