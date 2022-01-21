<template>
	<button @click="test">点击</button>
</template>

<script>
export default {
	data(){
		return{

		}
	},
	methods:{
		async test(){
			await this.$http.get("api/v1/file/2,01578d9194a0/",{responseType: 'arraybuffer'}).then((res)=>{
				let blob = new Blob([res.data]); // { type: "application/vnd.ms-excel" }
      let url = window.URL.createObjectURL(blob); // 创建一个临时的url指向blob对象
      // 创建url之后可以模拟对此文件对象的一系列操作，例如：预览、下载
      let a = document.createElement("a");
      a.href = url;
      a.download = "表格.docx";
      a.click();
      // 释放这个临时的对象url
      window.URL.revokeObjectURL(url);
			})
		}
	}
}
</script>

<style>

</style>