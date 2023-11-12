package domain

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Role", func() {
	var r Role
	Describe("ValidCreate 메소드는", func() {
		Describe("id 만 존재하면", func() {
			var roleID = 101
			BeforeEach(func() {
				r = NewRoleBuilder().Id(roleID).Build()
			})
			var err error
			BeforeEach(func() {
				err = r.ValidCreate()
			})
			It("에러를 반환한다", func() {
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("Update 메소드는", func() {
		var err error
		var updatedRole Role
		Describe("id 만 존재하면", func() {
			var roleID = 101
			BeforeEach(func() {
				r = NewRoleBuilder().Id(roleID).Build()
			})
			Context("[권한명]과 [권한설명]이 주어질때", func() {
				var (
					name        = "게시글관리"
					description = "모든 게시글을 관리할 수 있습니다."
				)
				BeforeEach(func() {
					updatedRole, err = r.Update(name, description)
				})
				It("변경된 [카페 권한]을 반환한다", func() {
					Expect(updatedRole).ToNot(BeNil())
					Expect(err).To(BeNil())
				})
			})
			Context("[권한설명]만 주어질경우", func() {
				var description = "마스터 권한입니다. 해당 카페내 모든걸 할수있습니다."
				BeforeEach(func() {
					updatedRole, err = r.Update("", description)
				})
				It("변경되지 않고, 에러를 반환한다", func() {
					Expect(updatedRole).ToNot(BeNil())
					Expect(err).ToNot(BeNil())
				})
			})
		})
		When("필드가 없으면", func() {
			BeforeEach(func() {
				r = NewRoleBuilder().Build()
			})
			Context("[권한명]과 [권한설명]이 주어질때", func() {
				var (
					name        = "게시글관리"
					description = "모든 게시글을 관리할 수 있습니다."
				)
				BeforeEach(func() {
					updatedRole, err = r.Update(name, description)
				})
				It("변경되지 않고 에러를 반환한다", func() {
					Expect(updatedRole).ToNot(BeNil())
					Expect(err).ToNot(BeNil())
				})
			})
		})
	})
})
