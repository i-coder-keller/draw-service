package service

import (
	"draw-service/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type ProjectList struct {
	ProjectName  string                `json:"projectName"`
	ProjectId    string                `json:"projectId"`
	ProjectInfo  string                `json:"projectInfo"`
	UpdateTime   int64                 `json:"updateTime"`
	Participants []*ProjectParticipant `json:"participants"`
	Owner        bool                  `json:"owner"`
}
type ProjectParticipant struct {
	Identity string `json:"identity"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name"`
}

// Projects 项目列表
func Projects(c *gin.Context) {
	result := make([]*ProjectList, 0)
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
				Identity: userInfo.Identity,
				Avatar:   userInfo.Avatar,
				Name:     userInfo.Nickname,
			})
		}
		// 根据项目Id查询项目信息
		projectInfo, _ := models.FindAllProjectByIdentity(ownerProject.ProjectIdentity)
		result = append(result, &ProjectList{
			ProjectName:  projectInfo.Name,
			ProjectId:    projectInfo.Identity,
			UpdateTime:   projectInfo.UpdatedAt,
			Participants: projectParticipantUserInfos,
			ProjectInfo:  projectInfo.Info,
			Owner:        true,
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
				Identity: userInfo.Identity,
				Avatar:   userInfo.Avatar,
				Name:     userInfo.Nickname,
			})
		}
		result = append(result, &ProjectList{
			ProjectName:  projectInfo.Name,
			ProjectId:    projectInfo.Identity,
			UpdateTime:   projectInfo.UpdatedAt,
			ProjectInfo:  projectInfo.Info,
			Participants: projectParticipantUserInfos,
			Owner:        false,
		})

	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
		"msg":  "查询成功",
	})
}

// CreateProject 创建项目
func CreateProject(c *gin.Context) {
	type projectInfo struct {
		Name string `json:"name"`
		Info string `json:"info"`
	}
	var request projectInfo
	c.ShouldBindJSON(&request)
	userIdentity := c.GetString("userIdentity")
	result, err := models.InsertProject(request.Name, request.Info)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "创建失败",
		})
		return
	}
	strId := result.(primitive.ObjectID).Hex()
	err = models.InsertProjectIdentityByOwnerIdentity(userIdentity, strId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
	})

}

// InviteOfProject 邀请人进入项目
func InviteOfProject(c *gin.Context) {
	type Project struct {
		ProjectId string   `json:"projectId"`
		Users     []string `json:"users"`
	}
	userIdentity := c.GetString("userIdentity")
	var request Project
	c.ShouldBindJSON(&request)
	_, err := models.FindAllProjectByIdentity(request.ProjectId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "暂无此项目",
		})
		return
	}
	_, err = models.ValidationProjectOfOwner(userIdentity, request.ProjectId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "暂无权限",
		})
		return
	}
	for _, userId := range request.Users {
		_, err := models.GetUserBasicByIdentity(userId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "邀请失败，暂无此用户",
			})
			return
		}
	}
	for _, userId := range request.Users {
		err = models.InsertProjectIdentityByParticipantIdentity(userId, request.ProjectId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "邀请失败",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "邀请成功",
	})
}
