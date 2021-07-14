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
					<a-button type="primary">新增</a-button>
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
				<a-button type="primary" style="margin-right:15px">编辑</a-button>
				<a-popconfirm title="确认删除？删除后无法恢复" ok-text="确认" cancel-text="取消" @confirm="delUser(data.ID)">
				  <a-button type="danger">删除</a-button>
				</a-popconfirm>
			</div>
			</template>
			</a-table>
		</a-card>
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
		}
	},
	created(){
		this.getUserList()

	},
	methods:{
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
		async delUser(id){
		  const { data: res } = await this.$http.delete(`api/admin/user/${id}/`)
			if(res.code != 200)return this.$message.error(res.msg)
			this.queryParam.pagenum = 1
			this.getUserList()
		}
	}

}
</script>

<style>
.actionSlot{
	display: flex;
	justify-content: center;
}

</style>