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
						<!-- <img v-if="artInfo.img" :src="baseurl+artInfo.img" style="width:120px;height:100px"> -->
						<img v-if="artInfo.img" :src="'http://localhost:8080/'+artInfo.img" style="width:120px;height:100px">
					</template>
				</a-upload>
			</a-form-model-item>
			<a-form-model-item label="内容" prop="content">
				<mavon-editor
            ref="md"
            placeholder="请输入文档内容..."
            :boxShadow="false"
            style="z-index:1;border: 1px solid #d9d9d9;height:50vh"
            v-model="artInfo.content"
            :toolbars="toolbars"
						@imgAdd="$imgAdd"
						@imgDel="$imgDel"
          />
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
export default {
	props:['id'],
	data(){
		return{
			// baseurl:'',
			artInfo:{
				id:0,
				title:'',
				cid:undefined,
				desc:'',
				content:'',
				img:''
			},
			Catelist:[],
			uploadUrl:Url+'/api/admin/file/',
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
			// markdown对象
			toolbars: {
        bold: true, // 粗体
        italic: true, // 斜体
        header: true, // 标题
        underline: true, // 下划线
        strikethrough: true, // 中划线
        mark: true, // 标记
        superscript: true, // 上角标
        subscript: true, // 下角标
        quote: true, // 引用
        ol: true, // 有序列表
        ul: true, // 无序列表
        link: true, // 链接
        imagelink: true, // 图片链接
        code: true, // code
        table: true, // 表格
        fullscreen: true, // 全屏编辑
        readmodel: true, // 沉浸式阅读
        htmlcode: true, // 展示html源码
        help: true, // 帮助
        /* 1.3.5 */
        undo: true, // 上一步
        redo: true, // 下一步
        trash: true, // 清空
        save: false, // 保存（触发events中的save事件）
        /* 1.4.2 */
        navigation: true, // 导航目录
        /* 2.1.8 */
        alignleft: true, // 左对齐
        aligncenter: true, // 居中
        alignright: true, // 右对齐
        /* 2.2.1 */
        subfield: true, // 单双栏模式
        preview: true // 预览
      }
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
			this.artInfo.img = res.data.img
			// 本来用于渲染图片前缀，通过v-if图片解决了
			// this.baseurl = "http://localhost:8081/"
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
          this.$router.push('/artlist')
          this.$message.success('添加文章成功')
        } else {
          const { data: res } = await this.$http.put(`/api/admin/article/${id}/`, this.artInfo)
          if (res.code !== 200) return this.$message.error(res.msg)
          this.$router.push('/artlist')
          this.$message.success('更新文章成功')
        }
			})
		},
		// 取消
		addCancel(){
			this.$refs.artInfoRef.resetFields()
		},
		// 上传图片方法
    async $imgAdd(pos, $file) {
			var formdata = new FormData();
			formdata.append('file', $file);
			const {data: res} = await this.$http.post("/api/admin/file/",formdata,{ 'Content-Type': 'multipart/form-data' })
			this.$refs.md.$imglst2Url([[pos, "http://localhost:8080/"+res.data]])
    },
		async $imgDel(pos){
			var fid = pos[0].split("/")
			const {data: res} = await this.$http.delete("/api/admin/file/",{params:{"fid":fid[fid.length-1]}})
			if (res.code !== 200) return this.$message.error(res.msg)
			return this.$message.success("删除成功")
		}
	}

}
</script>

<style>

</style>