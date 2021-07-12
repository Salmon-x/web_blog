<template>
	<a-button type="danger" @click="loginout">登出</a-button>
</template>

<script>
export default {
	methods:{
		loginout(){
			window.sessionStorage.clear('token')
			this.$router.push('/login')
		}
	}

}
</script>

<style>

</style>