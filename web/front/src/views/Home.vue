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
							class="page_drawer nav"
							style="width:20rem"
							absolute
							elevate-on-scroll
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
									<div class="ma-3">About Me:</div>
									<div class="ma-3">一个正在冉冉升起的博客</div>
									<del class="ma-5">虽然升起的速度很慢</del>
								</v-col>

								<v-divider color="indigo"></v-divider>

								<v-list nav dense>
									<v-list-item>
										<!-- <v-list-item-icon class="ma-3">
											<v-icon>{{'mdi-qqchat'}}</v-icon>
										</v-list-item-icon>
										<v-list-item-content class="grey-text">123123123</v-list-item-content> -->
										<v-list-item-content>第一版先这样啦</v-list-item-content>
									</v-list-item>
								</v-list>

							</v-card>
						</v-navigation-drawer>
						
					</v-col>
					
					<v-col>
						<v-sheet min-height="75vh" rounded="lg" :width="nav?'870':'1200'" style="transition: all 0.3s" >
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
.nav{
	display: inline-block;  
	position: fixed;  
	top: 23.2rem;  
	bottom: 0;  
	overflow: scroll;
	right: 0;  
	margin-bottom: 4.7rem;  
	transition: all 0.3s;
	z-index: 4;
}
</style>