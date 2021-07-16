<template>
	<div>
		<a-card>
			<a-row :gutter="20">
				<a-col :span="6">
					<a-input-search 
					v-model="queryParam.username"
					allowClear placeholder="请输入用户名" 
					enter-button 
					@search="getUserList" />
				</a-col>
				<a-col :span="4">
					<a-button type="primary" @click="AddUserVisible = true">新增</a-button>
				</a-col>
			</a-row>
			<a-table
				:columns='columns' 
				rowKey='ID'
				:pagination="pagination"
				:dataSource="userlist"
				bordered
				@change="handleTableChange"
			>
			<span slot="role" slot-scope="role">{{role == 1 ? "管理员":"订阅者"}}</span>
			<template slot="action" slot-scope="data">
				<div class="actionSlot">
				<a-button type="info" icon="info" style="margin-right:15px" @click="ChangePassword(data.ID)">修改密码</a-button>
				<a-button type="primary" icon="edit" style="margin-right:15px" @click="editUser(data.ID)">编辑</a-button>
				<a-popconfirm title="确认删除？删除后无法恢复" ok-text="确认" cancel-text="取消" @confirm="delUser(data.ID)">
				  <a-button icon="delete" type="danger">删除</a-button>
				</a-popconfirm>
			</div>
			</template>
			</a-table>
		</a-card>
		
		<!-- 新增用户对话框 -->
		<a-modal
      title="新增用户"
      :visible="AddUserVisible"
      @ok="AddUserOk"
      @cancel="addUserCanal"
			closable
    >
      <a-form-model :model="userinfo" :rules="userRules" ref="addUserRef">
				<a-form-model-item label="用户名" prop="username">
					<a-input v-model="userinfo.username"></a-input>
				</a-form-model-item>
				<a-form-model-item label="密码" prop="password">
					<a-input-password v-model="userinfo.password">

					</a-input-password>
				</a-form-model-item>
				<a-form-model-item label="确认密码" prop="checkpass">
					<a-input-password v-model="userinfo.checkpass">

					</a-input-password>
				</a-form-model-item>
				
			</a-form-model>
    </a-modal>

		<!-- 编辑用户 -->
		<a-modal
      title="编辑用户"
      :visible="editUserVisible"
      @ok="editUserOk"
			width="60%"
      @cancel="editUserCanal"
			closable
    >
      <a-form-model :model="userinfo" :rules="userRules" ref="addUserRef">
				<a-form-model-item label="用户名" prop="username">
					<a-input v-model="userinfo.username"></a-input>
				</a-form-model-item>

				<a-form-model-item label="是否为管理员">
          <a-switch
            :checked="IsAdmin"
            checked-children="是"
            un-checked-children="否"
            @change="adminChange"
          />
        </a-form-model-item>
				
			</a-form-model>
    </a-modal>

		<!-- 编辑密码 -->
		<a-modal
      closable
      title="修改密码"
      :visible="changePasswordVisible"
      width="60%"
      @ok="changePasswordOk"
      @cancel="changePasswordCancel"
      destroyOnClose
    >
      <a-form-model :model="userinfo" :rules="userRules" ref="addUserRef">
        <a-form-model-item has-feedback label="密码" prop="password">
          <a-input-password v-model="userinfo.password"></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label="确认密码" prop="checkpass">
          <a-input-password v-model="userinfo.checkpass"></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

	</div>
</template>

<script>
const columns = [
	{
    title: 'ID',
    dataIndex: 'ID',
    key: 'id',
    width: '10%',
		align:'center'
  },
	{
    title: '用户名',
    dataIndex: 'username',
    key: 'username',
    width: '20%',
		align:'center'
  },
	{
    title: '角色',
    dataIndex: 'role',
    key: 'role',
    width: '20%',
		scopedSlots:{customRender:'role'},
		align:'center'
  },
	{
    title: '操作',
    key: 'action',
    width: '30%',
		scopedSlots:{customRender:'action'},
		align:'center'
  },
]
export default {
	data(){
		return{
			columns,
			changePasswordVisible: false,
			// 分页
			pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
      },
			queryParam: {
				username:'',
        pagesize: 5,
        pagenum: 1,
      },
			userlist:[],
			AddUserVisible:false,
			editUserVisible:false,
			userinfo:{
				id:0,
				username:'',
				password:'',
				checkpass:'',
				role:2
			},
			userRules:{
				username: [
          {
            validator: (rule, value, callback) => {
              if (this.userinfo.username == '') {
                callback(new Error('请输入用户名'))
              }
              if ([...this.userinfo.username].length < 4 || [...this.userinfo.username].length > 12) {
                callback(new Error('用户名应当在4到12个字符之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.userinfo.password == '') {
                callback(new Error('请输入密码'))
              }
              if ([...this.userinfo.password].length < 6 || [...this.userinfo.password].length > 20) {
                callback(new Error('密码应当在6到20位之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],
				checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.userinfo.checkpass == '') {
                callback(new Error('请输入密码'))
              }
              if (this.userinfo.password !== this.userinfo.checkpass) {
                callback(new Error('密码不一致，请重新输入'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          },
        ],



			},
		}
	},
	created(){
		this.getUserList()
	},
	computed: {
		// 加载前布尔判断
    IsAdmin: function () {
      if (this.userinfo.role === 1) {
        return true
      } else {
        return false
      }
    }
	},
	methods:{
		// 查询用户
		async getUserList(){
			const { data:res } = await this.$http.get('api/v1/user/', {
				params:{
					username: this.queryParam.username,
					size: this.queryParam.pagesize,
					page: this.queryParam.pagenum,
					},
				})
			if(res.code != 200)return this.$message.error(res.msg)
			this.userlist = res.data
			this.pagination.total = res.total
		},
		// 分页
		handleTableChange(pagination, filters, sorter) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getUserList()
    },
		// 删除用户
		async delUser(id){
		  const { data: res } = await this.$http.delete(`api/admin/user/${id}/`)
			if(res.code != 200)return this.$message.error(res.msg)
			this.queryParam.pagenum = 1
			this.$message.success('删除成功')
			this.getUserList()
		},

		AddUserOk(){
			this.$refs.addUserRef.validate(async (vaild)=>{
				if(!vaild)return this.$message.error("格式不正确")
				const { data: res } = await this.$http.post("api/v1/user/",{
					"username":this.userinfo.username,
					"password":this.userinfo.password,
					"role": this.userinfo.role,
					})
				if (res.code != 200) return this.$message.error(res.msg)
				// 清空搜索框
				this.$refs.addUserRef.resetFields()
				// 关闭弹窗
        this.AddUserVisible = false
        this.$message.success('添加用户成功')
        this.getUserList()
			})
		},
		// role布尔值判断
		adminChange(checked){
      if (checked) {
        this.userinfo.role = 1
      } else {
        this.userinfo.role = 2
      }
		},
		// 取消的方法
		addUserCanal(){
			this.$refs.addUserRef.resetFields()
			this.AddUserVisible = false
		},
		// 编辑框获取用户信息
		async editUser(id){
			this.editUserVisible = true
			const {data: res } = await this.$http.get(`/api/v1/user/${id}/`)
			this.userinfo = res.data
			this.userinfo.id = id
		},
		// 编辑用户
		editUserOk(){
			this.$refs.addUserRef.validate(async (vaild)=>{
				if(!vaild)return this.$message.error("格式不正确")
				const { data: res } = await this.$http.put(`/api/admin/user/${this.userinfo.id}/`,{
					"username":this.userinfo.username,
					"role": this.userinfo.role,
					})
				if (res.code != 200) return this.$message.error(res.msg)
				// 清空搜索框
				this.$refs.addUserRef.resetFields()
				// 关闭弹窗
        this.editUserVisible = false
        this.$message.success('编辑用户成功')
        this.getUserList()
			})
		},
		// 取消编辑销毁数据
		editUserCanal(){
			this.$refs.addUserRef.resetFields()
			this.editUserVisible = false
		},
		// 修改密码获取信息
		ChangePassword(id){
			this.changePasswordVisible = true
      this.userinfo.id = id
		},
		// 修改密码
		changePasswordOk() {
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`/api/admin/editpass/${this.userinfo.id}/`,{
				"password":this.userinfo.password
			})
        if (res.code != 200) return this.$message.error(res.msg)
        this.changePasswordVisible = false
        this.$message.success('修改密码成功')
        this.getUserList()
      })
    },
		// 修改密码取消
		changePasswordCancel() {
      this.$refs.addUserRef.resetFields()
      this.changePasswordVisible = false
      this.$message.info('已取消')
    },


	}

}
</script>

<style>
.actionSlot{
	display: flex;
	justify-content: center;
}

</style>