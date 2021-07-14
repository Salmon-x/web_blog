<template>
    <div class="container">
        <div class="loginBox">
            <a-form-model ref="loginFormRef" :rules="rules" :model="formdata" class="loginForm">
                <a-form-model-item prop="username">
                    <a-input placeholder="账号" v-model="formdata.username">
                        <a-icon slot="prefix" type="user" style="color:rgba(0,0,0,.25)" />
                    </a-input>
                </a-form-model-item>
                <a-form-model-item prop="password">
									<!-- v-on:keyup.enter='login' 监听键盘enter，触发login方法 -->
                    <a-input placeholder="密码" v-on:keyup.enter='login' type="password" v-model="formdata.password">
                        <a-icon slot="prefix" type="lock" style="color:rgba(0,0,0,.25)" />
                    </a-input>
                </a-form-model-item>
                <a-form-model-item class="loginButton">
                    <a-button type="primary" style="margin:10px" @click="login">登陆</a-button>
                    <a-button type="info" style="margin:10px" @click="resetform">取消</a-button>
                </a-form-model-item>
            </a-form-model>
        </div>
    </div>
</template>

<script>
export default {
  data () {
		return {
      formdata: {
          username: '',
          password: ''
      },
	    rules: {
          username: [
              { required: true, message: '请输入管理员账号', trigger: 'blur' },
              { min: 4, max: 12, message: '用户名长度为4-12个字符', trigger: 'blur' },
          ],
					password: [
              { required: true, message: '请输入密码', trigger: 'blur' },
              { min: 6, max: 20, message: '密码长度为6-20个字符', trigger: 'blur' },
          ],
      }
    }
  },
	methods:{
		 resetform:function(){
			this.$refs.loginFormRef.resetFields()
		 },
		 login:function(){
			 this.$refs.loginFormRef.validate(async valid=>{
				 if(!valid) return this.$message.error("非法输入")
				 const { data:res } = await this.$http.post('api/v1/login/', this.formdata)
				 if(res.status != 200) return this.$message.error(res.message)
				 this.$message.info(res.message)
				 window.sessionStorage.setItem("token", res.token)
				 this.$router.push('admin/index')
			 })
		 }
	}
}

</script>

<style scoped>
    .container{
        height: 100%;
        background-color: aliceblue;
    }
    .loginBox{
        width: 450px;
        height: 300px;
        background-color: white;
        /* 绝对定位 */
				position: absolute;
        top: 50%;
        left: 70%;
				/* 旋转 */
        transform: translate(-50%, -50%);
				/* 圆角 */
        border-radius: 9px;
    }
    .loginForm{
        width: 100%;
        position: absolute;
				/* 距离底部边缘 */
        bottom: 10%;
        /* 设置内边距 */
				padding: 0 20px;
        box-sizing: border-box;
    }
    .loginButton{
        display: flex;
				/* 中央对齐 */
        justify-content: flex-end;
    }

</style>>
