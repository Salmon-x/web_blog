<template>
<div>
	<v-app-bar app color="#FFF">
		<v-btn
				fab
				style="margin-right: 16px"
				text
				@click="handleMiniMenu"
		>
				<v-icon v-if="!miniVariant">mdi-view-headline</v-icon>
				<v-icon v-else>mdi-close</v-icon>
		</v-btn>

		<!-- 按钮和搜索框 -->
		<v-container class="py-0 file-height">
			<v-btn text @click="$router.push('/').catch((err)=>err)">
			<v-icon>mdi-bank</v-icon>
			&nbsp;主页</v-btn>
			<!-- <v-btn v-for="item in cateList" :key="item.id" text>{{item.name}}</v-btn> -->
			<v-btn text @click="$router.push('/category/').catch((err)=>err)">
			<v-icon>mdi-book-variant</v-icon>
			&nbsp;类别</v-btn>
			<v-btn text @click="$router.push('/tag/').catch((err)=>err)">
			<v-icon>mdi-tag</v-icon>
			&nbsp;标签</v-btn>
		</v-container>

	<v-responsive max-width="260" color="white">
		<v-text-field dense flat hide-details rounded solo-inverted></v-text-field>
	</v-responsive>

	</v-app-bar>
</div>
	
</template>

<script>
export default {
	data(){
		return{
			cateList:[],
			miniVariant: false,
		}
	},
	created(){
		this.getcateList()
	},
	methods:{
		// 获取分类
		async getcateList(){
			const { data:res } = await this.$http.get('api/v1/category/', {
				params:{
					},
				})
			if(res.code != 200)return this.$message.error(res.msg)
			this.cateList = res.data
			// this.pagination.total = res.total

		},
		handleMiniMenu(){
			this.miniVariant = !this.miniVariant;
			this.$emit("miniVariant",this.miniVariant)
		}
	}

}
</script>

<style>

</style>