<template>
	<div>
		<a-card>
			<a-row :gutter="20">
				<a-col :span="6">
					<a-input-search placeholder="请输入用户名" enter-button />
				</a-col>
				<a-col :span="4">
					<a-button type="primary">新增</a-button>
				</a-col>
			</a-row>
			<a-table 
				:columns='columns' 
				rowKey='ID' 
				:pagination="PaginationOption" 
				:dataSource="userlist"
				bordered
			>
			<span slot="role" slot-scope="role">{{role == 1 ? "管理员":"订阅者"}}</span>
			<template slot="action">
				<div class="actionSlot">
				<a-button type="primary" style="margin-right:15px">编辑</a-button>
				<a-button type="danger">删除</a-button>
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
			PaginationOption:{
				pageSizeOptions:['5','10','20'],
				defaultCurrent:1,
				defaultPageSize:5,
				total:0,
				showSizeChanger:true,
				showTotal:(total)=>`共${total}条`,
				onChage:(page, pageSize)=>{
					this.PaginationOption.defaultCurrent = page
					this.PaginationOption.defaultPageSize = pageSize
					this.getUserList()
				},
				onshowSizeChange:(current, size)=>{
					this.PaginationOption.defaultCurrent = current
					this.PaginationOption.defaultPageSize = size
					this.getUserList()
				}

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
					size:this.PaginationOption.defaultPageSize, 
					page:this.PaginationOption.defaultCurrent
					},
				})
				if(res.code != 200)return this.$message.error(res.msg)
				this.userlist = res.data
				this.PaginationOption.total = res.total
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