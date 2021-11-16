<template>
	<div>
		<div class="d-flex justify-center ps-3 text-h2 font-weight-bold">{{artInfo.title}}</div>
		<v-divider class="pa-3 ma-3"></v-divider>
		<v-alert color='grey' dark border='left' outlined class="ma-4" elevation="2">
			<v-icon>mdi-calendar</v-icon>
						<span>Publish Date：{{artInfo.CreatedAt | dataformat("YYYY-MM-DD")}}</span>
						<v-icon>mdi-calendar-clock</v-icon>
						<span>Update Date：{{artInfo.UpdatedAt | dataformat("YYYY-MM-DD")}}</span>
		</v-alert>
		<div v-html="artInfo.content" class="content ma-5 pa-3 text-justify"></div>
	</div>
</template>

<script>
export default {
	props:['id'],
	data(){
		return{
			artInfo:{},
			
		}
	},
	created(){
		this.getArtInfo()
	},
	methods:{
		async getArtInfo(){
			const { data: res } = await this.$http.get(`api/v1/frontarticle/${this.id}/`)
			this.artInfo = res.data
		},
	}

}
</script>

<style lang="scss" scoped>
.content >>> img,span,p{
	width: 90%;
}
</style>