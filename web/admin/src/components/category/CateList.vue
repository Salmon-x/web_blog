<template>
	<div>
		<a-card>
			<a-row :gutter="20">
				<a-col :span="4">
					<a-button type="primary" @click="AddcateVisible = true">新增分类</a-button>
				</a-col>
			</a-row>
			<a-table
				:columns='columns' 
				rowKey='id'
				:pagination="pagination"
				:dataSource="catelist"
				bordered
				@change="handleTableChange"
			>
			<template slot="action" slot-scope="data">
				<div class="actionSlot">
				<a-button type="primary" icon="edit" style="margin-right:15px" @click="editcate(data.id)">编辑</a-button>
				<a-popconfirm title="确认删除？删除后无法恢复" ok-text="确认" cancel-text="取消" @confirm="delcate(data.id)">
				  <a-button icon="delete" type="danger">删除</a-button>
				</a-popconfirm>
			</div>
			</template>
			</a-table>
		</a-card>
		
		<!-- 新增分类对话框 -->
		<a-modal
      title="新增分类"
      :visible="AddcateVisible"
      @ok="AddcateOk"
      @cancel="addcateCanal"
			closable
    >
      <a-form-model :model="cateinfo" ref="addcateRef">
				<a-form-model-item label="分类名" prop="name">
					<a-input v-model="cateinfo.name"></a-input>
				</a-form-model-item>
			</a-form-model>
    </a-modal>

		<!-- 编辑分类 -->
		<a-modal
      title="编辑分类"
      :visible="editcateVisible"
      @ok="editcateOk"
			width="60%"
      @cancel="editcateCanal"
			closable
    >
      <a-form-model :model="cateinfo" ref="addcateRef">
				<a-form-model-item label="分类名" prop="name">
					<a-input v-model="cateinfo.name"></a-input>
				</a-form-model-item>
				
			</a-form-model>
    </a-modal>
	</div>
</template>

<script>
const columns = [
	{
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
    width: '10%',
		align:'center'
  },
	{
    title: '分类名',
    dataIndex: 'name',
    key: 'name',
    width: '20%',
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
				name:'',
        pagesize: 5,
        pagenum: 1,
      },
			catelist:[],
			AddcateVisible:false,
			editcateVisible:false,
			cateinfo:{
				id:0,
				name:'',
			},
		}
	},
	created(){
		this.getcateList()
	},
	methods:{
		// 查询分类
		async getcateList(){
			const { data:res } = await this.$http.get('api/v1/category/', {
				params:{
					size: this.queryParam.pagesize,
					page: this.queryParam.pagenum,
					},
				})
			if(res.code != 200)return this.$message.error(res.msg)
			this.catelist = res.data
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
      this.getcateList()
    },
		// 删除分类
		async delcate(id){
		  const { data: res } = await this.$http.delete(`api/admin/category/${id}/`)
			if(res.code != 200)return this.$message.error(res.msg)
			this.$message.success('删除成功')
			this.queryParam.pagenum = 1
			this.getcateList()
		},

		async AddcateOk(){
				const { data: res } = await this.$http.post("api/admin/category/",{
					"name":this.cateinfo.name,
					})
				if (res.code != 200) return this.$message.error(res.msg)
				// 清空搜索框
				this.$refs.addcateRef.resetFields()
				// 关闭弹窗
        this.AddcateVisible = false
        this.$message.success('添加成功')
        this.getcateList()
		},
		// 取消的方法
		addcateCanal(){
			this.$refs.addcateRef.resetFields()
			this.AddcateVisible = false
		},
		// 编辑框获取分类信息
		async editcate(id){
			this.editcateVisible = true
			const {data: res } = await this.$http.get(`/api/v1/category/${id}/`)
			this.cateinfo = res.data
			this.cateinfo.id = id
		},
		// 编辑分类
		async editcateOk(){
				const { data: res } = await this.$http.put(`/api/admin/category/${this.cateinfo.id}/`,{
					"name":this.cateinfo.name,
					})
				if (res.code != 200) return this.$message.error(res.msg)
				// 清空搜索框
				this.$refs.addcateRef.resetFields()
				// 关闭弹窗
        this.editcateVisible = false
        this.$message.success('编辑成功')
        this.getcateList()
		},
		// 取消编辑销毁数据
		editcateCanal(){
			this.$refs.addcateRef.resetFields()
			this.editcateVisible = false
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