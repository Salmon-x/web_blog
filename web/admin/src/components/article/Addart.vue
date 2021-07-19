<template>
	<div>
		<a-card>
			<h3>{{id?'编辑文章':'新增文章'}}</h3>
		<a-form-model :model="artInfo" ref="artInfoRef" :rules="artInfoRules" :hideRequiredMark="true">
			<a-form-model-item label="文章标题" prop="title">
				<a-input style="width:300px" v-model="artInfo.title"></a-input>
			</a-form-model-item>
			<a-form-model-item 
			style="width:200px"
			prop="cid"
			label="文章分类">
				<a-select v-model="artInfo.cid" placeholder="请选择分类" @change="cateChange">
					<a-select-option v-for="item in Catelist" :key="item.id" :value="item.id">{{item.name}}</a-select-option>
				</a-select>
			</a-form-model-item>
			<a-form-model-item label="文章描述" prop="desc">
				<a-input type="textarea" v-model="artInfo.desc"></a-input>
			</a-form-model-item>
			<a-form-model-item label="缩略图" prop="img"> 
				<a-upload
					name="file"
					:multiple="false"
					:action="uploadUrl"
					:headers="headers"
					@change="upChange"
					listType="picture"
				>
					<a-button> <a-icon type="upload" /> 上传图片 </a-button>
					<br/>
					<template v-if="id">
						<img :src="artInfo.img" style="width:120px;height:100px">
					</template>
				</a-upload>
			</a-form-model-item>
			<a-form-model-item label="内容" prop="content">
				<Editor v-model="artInfo.content"></Editor>
			</a-form-model-item>
			<a-form-model-item>
				<a-button type="primary" 
				style="margin-right:15px" 
				@click="artOk(artInfo.id)"
				>
				{{artInfo.id?"更新":'提交'}}
				</a-button>
				<a-button type="danger" @click="addCancel">取消</a-button>
			</a-form-model-item>
		</a-form-model>
		</a-card>
	</div>
</template>

<script>
import {Url} from "../../plugins/http"
import Editor from "../editor/index"
export default {
	components:{Editor},
	props:['id'],
	data(){
		return{
			artInfo:{
				id:0,
				title:'',
				cid:undefined,
				desc:'',
				content:'',
				img:''
			},
			Catelist:[],
			uploadUrl:Url+'/api/admin/upload/',
			headers:{},
			artInfoRules: {
        title: [{ required: true, message: '请输入文章标题', trigger: 'change' }],
        cid: [{ required: true, message: '请选择文章分类', trigger: 'change' }],
        desc: [
          { required: true, message: '请输入文章描述', trigger: 'change' },
          { max: 120, message: '描述最多可写120个字符', trigger: 'change' },
        ],
        content: [{ required: true, message: '请输入文章内容', trigger: 'change' }],
      },
		}
	},
	created(){
		this.getCateList()
		this.headers = {Authorization:`Bearer ${window.sessionStorage.getItem('token')}`}
		if(this.id){
			this.getArtInfo(this.id)
		}
	},
	methods:{
		// 查询文章信息
		async getArtInfo(id){
			const {data: res} = await this.$http.get(`api/v1/article/${id}/`)
			if(res.code != 200) return this.$$message.error(res.msg)
			this.artInfo = res.data
			this.artInfo.id = res.data.ID
		},
		// 获取分类
    async getCateList() {
      const { data: res } = await this.$http.get('api/v1/category/',{params: {
          size: 100,
          page: 1,
        },})
      if (res.code !== 200) return this.$message.error(res.msg)
      this.Catelist = res.data
    },
		cateChange(val){
			this.artInfo.cid = val
		},
		upChange(info){
			if (info.file.status !== 'uploading') {
        
      }
      if (info.file.status === 'done') {
        this.$message.success('上传成功');
				const imgUrl = info.file.response.data
				this.artInfo.img = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error(`上传失败`);
      }
		},
		// 提交
		artOk(id){
			this.$refs.artInfoRef.validate(async (valid) => {
				if (!valid) return this.$message.error('参数验证未通过，请按要求录入文章内容')
				if (id === 0) {
          const { data: res } = await this.$http.post('/api/admin/article/', this.artInfo)
          if (res.code !== 200) return this.$message.error(res.msg)
          this.$router.push('/admin/artlist')
          this.$message.success('添加文章成功')
        } else {
          const { data: res } = await this.$http.put(`/api/admin/article/${id}/`, this.artInfo)
          if (res.code !== 200) return this.$message.error(res.msg)
          this.$router.push('/admin/artlist')
          this.$message.success('更新文章成功')
        }
			})
		},
		// 取消
		addCancel(){
			this.$refs.artInfoRef.resetFields()
		}
	}

}
</script>

<style>

</style>