<template>
  <v-app app>
		<TopBar
		@miniVariant="miniVariant"
		></TopBar>
		<v-main class="grey lighten-3">
			<v-container>
				<v-row>
					<v-col>
						<!-- nav -->
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
					</v-col>
					<v-col>
						<v-sheet min-height="75vh" rounded="lg" :width="nav?'900':'1200'" style="transition: all 0.3s">
							<!-- 将路由匹配到的组件显示在这里 -->
							<router-view></router-view>
						</v-sheet>
					</v-col>
				</v-row>
			</v-container>
		</v-main>

		<Footer></Footer>
	</v-app>
</template>

<script>
import TopBar from '../components/TopBar.vue'
import Footer from '../components/Footer.vue'
// import Nav from '../components/Nav.vue'
// import ArticleList from '../components/ArticleList.vue'
  export default {
		components:{TopBar,Footer},
		data(){
			return{
				nav:false,
				wks:"",
			}
		},
		created(){
			this.getwks()
		},
		methods:{
			// 子组件更改值，显示左侧人物介绍
			miniVariant(data){
				this.nav = data
			},
			async getwks(){
					const { data:res } = await this.$http.get("api/v1/wks/")
					if(res.code != 200)return this.$message.error(res.msg)
					var y = Math.round(Math.random()*((res.data.length-1)-0)+0)
					this.wks = res.data[y].title
			},
		}
  }
</script>

<style scoped>
</style>