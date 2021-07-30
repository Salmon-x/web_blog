<template>
	<v-navigation-drawer
							v-model="nav"
							disable-resize-watcher
							mini-variant-width="74"
							class="page_drawer"
							absolute
							elevate-on-scroll
							style="transition: all 0.3s;z-index: 4"
							>
							<v-card class="mx-auto" max-width="300">
								<v-img src="../assets/nav_bg.jpeg">
									<v-card-title>
										<v-col align="center">
											<v-avatar size="130" color="grey">
												<img src="../assets/头像.jpg">
											</v-avatar>

											<div class="ma-4 black--text">{{this.wks}}</div>
										</v-col>
									</v-card-title>
									<v-divider></v-divider>
								</v-img>
								<v-col>
									<div class="ma-3">About</div>
									<div class="ma-3">友情链接</div>
								</v-col>

								<v-divider color="indigo"></v-divider>

								<v-list nav dense>
									<v-list-item>
										<v-list-item-icon class="ma-3">
											<v-icon>{{'mdi-qqchat'}}</v-icon>
										</v-list-item-icon>
										<v-list-item-content class="grey-text">123123123</v-list-item-content>
									</v-list-item>
									<v-list-item>
										<v-list-item-icon class="ma-3">
											<v-icon>{{'mdi-qqchat'}}</v-icon>
										</v-list-item-icon>
										<v-list-item-content class="grey-text">123123123</v-list-item-content>
									</v-list-item>
									<v-list-item>
										<v-list-item-icon class="ma-3">
											<v-icon>{{'mdi-qqchat'}}</v-icon>
										</v-list-item-icon>
										<v-list-item-content class="grey-text">123123123</v-list-item-content>
									</v-list-item>
									<v-list-item>
										<v-list-item-icon class="ma-3">
											<v-icon>{{'mdi-qqchat'}}</v-icon>
										</v-list-item-icon>
										<v-list-item-content class="grey-text">123123123</v-list-item-content>
									</v-list-item>
									<v-list-item>
										<v-list-item-icon class="ma-3">
											<v-icon>{{'mdi-qqchat'}}</v-icon>
										</v-list-item-icon>
										<v-list-item-content class="grey-text">123123123</v-list-item-content>
									</v-list-item>
									<v-list-item>
										<v-list-item-icon class="ma-3">
											<v-icon>{{'mdi-qqchat'}}</v-icon>
										</v-list-item-icon>
										<v-list-item-content class="grey-text">123123123</v-list-item-content>
									</v-list-item>
									<v-list-item>
										<v-list-item-icon class="ma-3">
											<v-icon>{{'mdi-qqchat'}}</v-icon>
										</v-list-item-icon>
										<v-list-item-content class="grey-text">123123123</v-list-item-content>
									</v-list-item>
								</v-list>

							</v-card>
	</v-navigation-drawer>
</template>

<script>
export default {
	data(){
		return{
			wks:"",
		}
	},
	created(){
		this.getwks()
	},
	methods:{
		async getwks(){
			const { data:res } = await this.$http.get("api/v1/wks/")
			if(res.code != 200)return this.$message.error(res.msg)
			var y = Math.round(Math.random()*((res.data.length-1)-0)+0)
			this.wks = res.data[y].title
		}
	}

}
</script>

<style>

</style>