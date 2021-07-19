<template>
		<a-layout class="container">
				<Nav></Nav>
      <a-layout>
        
				<a-layout-header class="headerBtn">
					<Header></Header>	
				</a-layout-header>
        
				<a-layout-content>
					<router-view :key="$route.path"></router-view>
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

</style>