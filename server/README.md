## 介绍
该项目使用Golang进行构建
以下是将您提供的数据整理成的表格形式：

---

### 1. **用户数据表 (`t_user`)**
| id | type | username | email          | email_verified | nickname   | avatar | gender | birthday | background_image | password (加密)                                                                 | home_page | description               | score | status | topic_count | comment_count | follow_count | fans_count | roles  | forbidden_end_time | create_time           | update_time           |
|----|------|----------|----------------|----------------|------------|--------|--------|----------|------------------|--------------------------------------------------------------------------------|-----------|---------------------------|-------|--------|-------------|---------------|--------------|------------|--------|--------------------|-----------------------|-----------------------|
| 1  | 0    | admin    | a@example.com  | 0              | duck站长   |        |        | NULL     | NULL             | $2a$10$ofA39bAFMpYpIX/Xiz7jtOMH9JnPvYfPRlzHXqAtLPFpbE/cLdjmS | NULL      | 轻轻地我走了，正如我轻轻的来。 | 0     | 0      | 0           | 0             | 0            | 0          | owner  | 0                  | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |

---

### 2. **角色数据表 (`t_role`)**
| id | type | name       | code  | sort_no | remark     | status | create_time           | update_time           |
|----|------|------------|-------|---------|------------|--------|-----------------------|-----------------------|
| 1  | 0    | 超级管理员 | owner | 0       | 超级管理员 | 0      | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 2  | 0    | 管理员     | admin | 1       | 管理员     | 0      | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |

---

### 3. **用户角色数据表 (`t_user_role`)**
| id | user_id | role_id | create_time           |
|----|---------|---------|-----------------------|
| 1  | 1       | 1       | UNIX_TIMESTAMP * 1000 |

---

### 4. **节点数据表 (`t_topic_node`)**
| id | name     | description | logo | sort_no | status | create_time           |
|----|----------|-------------|------|---------|--------|-----------------------|
| 1  | 默认节点 |             | NULL | 0       | 0      | UNIX_TIMESTAMP * 1000 |

---

### 5. **系统配置数据表 (`t_sys_config`)**
| id | key               | value                                                                 | name               | description               | create_time           | update_time           |
|----|-------------------|-----------------------------------------------------------------------|--------------------|---------------------------|-----------------------|-----------------------|
| 1  | siteTitle         | duck演示站                                                           | 站点标题           | 站点标题                  | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 2  | siteDescription   | duck，基于Go语言的开源社区系统                                       | 站点描述           | 站点描述                  | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 3  | siteKeywords      | ["duck"]                                                             | 站点关键字         | 站点关键字                | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 4  | siteNavs          | [{"title":"首页","url":"/"},{"title":"话题","url":"/topics"},...]     | 站点导航           | 站点导航                  | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 5  | defaultNodeId     | 1                                                                     | 默认节点           | 默认节点                  | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 6  | tokenExpireDays   | 365                                                                   | 用户登录有效期(天) | 用户登录有效期(天)        | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 7  | scoreConfig       | {"postTopicScore":1,"postCommentScore":1,"checkInScore":1}            | 积分配置           | 积分配置                  | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 8  | urlRedirect       | false                                                                 |                    |                           | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 9  | enableHideContent | false                                                                 |                    |                           | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 10 | siteLogo          |                                                                       |                    |                           | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 11 | siteNotification  |                                                                       |                    |                           | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 12 | recommendTags     |                                                                       |                    |                           | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 13 | modules           | {"tweet":true,"topic":true,"article":true}                            |                    |                           | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |

---

### 6. **菜单数据表 (`t_menu`)**
| id | parent_id | type  | name           | title      | icon          | path            | component            | sort_no | status | create_time           | update_time           |
|----|-----------|-------|----------------|------------|---------------|-----------------|----------------------|---------|--------|-----------------------|-----------------------|
| 1  | 0         | menu  | Dashboard      | 仪表盘     | icon-dashboard| /dashboard      | dashboard/index      | 0       | 0      | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 2  | 0         | menu  | User           | 用户管理   | icon-user     | /user           | user/index           | 1       | 0      | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 4  | 0         | menu  | Permission     | 权限管理   | icon-lock     |                 | NULL                 | 9       | 0      | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| 5  | 4         | menu  | Role           | 角色管理   |               | /permission/role| system/role/index    | 10      | 0      | UNIX_TIMESTAMP * 1000 | UNIX_TIMESTAMP * 1000 |
| ...| ...       | ...   | ...            | ...        | ...           | ...             | ...                  | ...     | ...    | ...                   | ...                   |

---

### 7. **角色菜单数据表 (`t_role_menu`)**
| id | role_id | menu_id | create_time           |
|----|---------|---------|-----------------------|
| 1  | 1       | 1       | UNIX_TIMESTAMP * 1000 |
| 2  | 1       | 2       | UNIX_TIMESTAMP * 1000 |
| 3  | 1       | 4       | UNIX_TIMESTAMP * 1000 |
| ...| ...     | ...     | ...                   |

---

如果需要进一步调整或补充，请告诉我！