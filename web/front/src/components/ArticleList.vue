<template>
	<v-col>
		<!-- 内容卡片 -->
		<v-card class="ma-3" v-for="item in Artlist" :key="item.id" link @click="$router.push(`details/${item.ID}/`)">
			<v-row no-gutters>
				<!-- <v-col class="d-flex justify-center align-center mx-3" cols="1"> -->
					<!-- style="opacity: .7" -->
					<v-img
					width="100%"
					height="440px"
					alt='图片正在加载'
					:src="seaweedUrl+item.img">
					<v-card-title style="justify-content: center;font-size: 60; color:black">
						{{item.title}}
					</v-card-title>
					</v-img>
				<v-col>



					<v-card-subtitle v-text="item.desc" style="text-align: center"></v-card-subtitle>
					<v-card-text style="text-align: center">
						<v-icon>mdi-calendar</v-icon>
						<span>{{item.CreatedAt | dataformat("YYYY-MM-DD")}}</span>
						<span class="post-meta-divider" style="margin: 0 .5rem">-</span>
						<v-icon>mdi-calendar-clock</v-icon>
						<span>{{item.UpdatedAt | dataformat("YYYY-MM-DD")}}</span>
					</v-card-text>
					<!-- 分割线 -->
					<v-divider></v-divider>
					<v-card-text>
							<v-chip style="text-align: left" color="pink" class="white--text">{{item.Category.name}}</v-chip>

						<span style="float: right">
							<v-icon>mdi-tag-multiple</v-icon>
							<span>标签</span>
						</span>
					</v-card-text>

				</v-col>
			</v-row>
		</v-card>

		<!-- 分页 -->
		<div class="text-center">
			<v-pagination
			v-model="queryParam.pagenum"
			:length="Math.ceil(this.total/queryParam.pagesize)"
			@input="getArtList()"
			></v-pagination>
		</div>
	</v-col>
</template>

<script>
export default {
	data(){
		return{
			seaweedUrl:'http://localhost:8080/',
			Artlist:[],
			queryParam:{
				title:'',
				pagesize:5,
				pagenum:1,
			},
			total:0,
		}
	},
	created(){
		this.getArtList()
	},
	methods:{
		// 获取文章列表
    async getArtList() {
      const { data: res } = await this.$http.get('api/v1/article/', {
        params: {
          title: this.queryParam.title,
          size: this.queryParam.pagesize,
          page: this.queryParam.pagenum,
        },
      })
			this.Artlist = res.data
      this.total = res.total

		window.pageYOffset = 0;
		document.documentElement.scrollTop = 0
		document.body.scrollTop = 0
    },
	}
}
</script>

<style>


</style>