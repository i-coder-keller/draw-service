package service

import (
	"draw-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProjectList struct {
	projectName  string                `json:"projectName"`
	projectId    string                `json:"projectId"`
	updateTime   int64                 `json:"updateTime"`
	participants []*ProjectParticipant `json:"participants"`
	owner        bool                  `json:"owner"`
}
type ProjectParticipant struct {
	identity string `json:"identity"`
	avatar   string `json:"avatar"`
	name     string `json:"name"`
}

func Projects(c *gin.Context) {
	var result []*ProjectList
	userIdentity := c.GetString("userIdentity")
	// 查询当前用户创建的项目
	ownerProjects, err := models.FindAllProjectIdentityByOwnerIdentity(userIdentity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "查询失败",
		})
		return
	}
	for _, ownerProject := range ownerProjects {
		var projectParticipantUserInfos []*ProjectParticipant
		// 根据创建的项目ID查询项目参与人
		projectParticipants, _ := models.FindAllProjectIdentityByParticipantIdentity(ownerProject.ProjectIdentity)
		for _, projectParticipant := range projectParticipants {
			// 根据项目参与人查询查询参与人信息
			userInfo, _ := models.GetUserBasicByIdentity(projectParticipant.ParticipantIdentity)
			projectParticipantUserInfos = append(projectParticipantUserInfos, &ProjectParticipant{
				identity: userInfo.Identity,
				avatar:   userInfo.Avatar,
				name:     userInfo.Nickname,
			})
		}
		// 根据项目Id查询项目信息
		projectInfo, _ := models.FindAllProjectByIdentity(ownerProject.ProjectIdentity)
		result = append(result, &ProjectList{
			projectName:  projectInfo.Name,
			projectId:    projectInfo.Identity,
			updateTime:   projectInfo.UpdatedAt,
			participants: projectParticipantUserInfos,
			owner:        true,
		})
	}

	// 查询当前用户参与的项目的项目Id
	participantProjects, err := models.FindAllParticipantIdentityByProjectIdentity(userIdentity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "查询失败",
		})
		return
	}
	for _, participantProject := range participantProjects {
		var projectParticipantUserInfos []*ProjectParticipant
		// 根据参与的项目Id查询项目详细信息
		projectInfo, _ := models.FindAllProjectByIdentity(participantProject.ParticipantIdentity)
		// 根据项目ID查询项目所有的参与人Id
		participantIdentitys, _ := models.FindAllProjectIdentityByParticipantIdentity(projectInfo.Identity)
		for _, projectData := range participantIdentitys {
			// 根据参与人ID查询参与人详细信息
			userInfo, _ := models.GetUserBasicByIdentity(projectData.ParticipantIdentity)
			projectParticipantUserInfos = append(projectParticipantUserInfos, &ProjectParticipant{
				identity: userInfo.Identity,
				avatar:   userInfo.Avatar,
				name:     userInfo.Nickname,
			})
		}
		result = append(result, &ProjectList{
			projectName:  projectInfo.Name,
			projectId:    projectInfo.Identity,
			updateTime:   projectInfo.UpdatedAt,
			participants: projectParticipantUserInfos,
			owner:        false,
		})

	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
		"msg":  "查询成功",
	})
}
