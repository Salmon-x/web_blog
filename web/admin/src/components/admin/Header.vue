<template>
<a-popconfirm title="确认退出？" ok-text="是" cancel-text="否" @confirm="loginout">
		<a-button type="danger">登出</a-button>
		<!-- <a href="#">Delete</a> -->
</a-popconfirm>
</template>

<script>
export default {
	methods:{
		loginout(){
			window.sessionStorage.clear('token')
			this.$router.push('/login')
			this.$message.success("感谢您的使用")
		}
	}

}
</script>

<style>

</style>