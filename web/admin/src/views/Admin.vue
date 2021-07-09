<template>
		<a-layout class="container">
      <a-layout-sider>
				<div class="log">
					<span>Salmon blog</span>
				</div>
				<Nav></Nav>
			</a-layout-sider>
      <a-layout>
        
				<a-layout-header class="headerBtn">
					<a-button type="danger" @click="loginout">登出</a-button>
				</a-layout-header>
        
				<a-layout-content>
					<router-view></router-view>
				</a-layout-content>

        <a-layout-footer>
					<Footer></Footer>
				</a-layout-footer>
      </a-layout>
    </a-layout>
</template>

<script>
import Footer from '../components/admin/Footer.vue'
import Nav from '../components/admin/Nav.vue'
import Header from '../components/admin/Header.vue'
export default {
	components:{Footer, Nav, Header},
	methods:{
		loginout(){
			window.sessionStorage.clear('token')
			this.$router.push('/login')
		}
	}

}
</script>

<style>
  .container {
		height: 100%;
	}
	.headerBtn{
		display: flex;
		justify-content: flex-end;
		align-items: center;
	}
	.log{
		height: 32px;
		margin: 16px;
		background-color: aliceblue;
		display: flex;
		justify-content: center;
		align-items: center;
		font-size: 17px;
	}

</style>