<template>
  <v-app app>
		<TopBar
		@miniVariant="miniVariant"
		></TopBar>
		<v-main class="grey lighten-3">
			<v-container>
				<v-row>
					<v-col cols="3" v-show="nav">
						<Nav></Nav>
					</v-col>
					<v-col>
						<v-sheet min-height="80vh" rounded="lg">
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
import Nav from '../components/Nav.vue'
// import ArticleList from '../components/ArticleList.vue'
  export default {
		components:{TopBar,Footer,Nav},
		data(){
			return{
				nav:false,
			}
		},
		methods:{
			// 子组件更改值，显示左侧人物介绍
			miniVariant(data){
				this.nav = data
			}
		}
  }
</script>

<style scoped>
</style>