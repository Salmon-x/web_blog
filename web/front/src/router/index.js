import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import ArticleList from "../components/ArticleList.vue"
import Category from "../components/Category.vue"
import Tags from "../components/Tags.vue"
import Details from "../components/Details.vue"

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
		children:[
			{"path":'/',component:ArticleList, meta:{'title':"天听"}},
			{"path":'/category/',component:Category, meta:{'title':"天听"}},
			{"path":'/tag/',component:Tags, meta:{'title':"天听"}},
			{"path":'/details/:id/',component:Details, meta:{'title':"天听"}, props:true}
		]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to,from ,next)=>{
	if(to.meta.title){
		document.title = to.meta.title
	}
	next()
})

export default router
