<template>
    <div class="login-container">
        <el-form :model="ruleForm2" :rules="rules2"
         status-icon
         ref="ruleForm2" 
         label-position="left" 
         label-width="0px" 
         class="demo-ruleForm login-page">
            <h3 class="ghd-title">GreenHouseDoc 用户登录</h3>
            <el-form-item prop="email">
                <el-input type="text" 
                    v-model="ruleForm2.email" 
                    auto-complete="off" 
                    placeholder="邮箱"
                ></el-input>
            </el-form-item>

            <el-form-item prop="password">
                <el-input type="password" 
                    v-model="ruleForm2.password" 
                    auto-complete="off" 
                    placeholder="密码"
                ></el-input>
            </el-form-item>

            <el-form-item>
                <el-input type="text"
                    v-model="ruleForm2.verifycode" 
                    auto-complete="off" 
                    placeholder="验证码" style="width:320px;"></el-input>
                <img style="float:right; height: 32px; line-height: 32px;" src="http://wiki.evente.cn/captcha" />
            </el-form-item>

            <el-form-item style="margin: 0;">
                <el-checkbox v-model="checked" class="rememberme">记住密码</el-checkbox>
                <router-link class="el-link-btn" style="float: right; color: #666;" to="forget">忘记密码？</router-link>
            </el-form-item>

            <el-form-item style="width:100%; margin-bottom: 0;">
                <el-button type="primary" style="width:100%;" @click="handleSubmit" :loading="logining">登录</el-button>
            </el-form-item>

            <el-form-item>
                还有没帐号？<router-link class="el-link-btn" to="reg">立即注册</router-link>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
export default {
    data(){
        return {
            logining: false,
            ruleForm2: {
                email: 'admin@gmail.com',
                password: '',
                verifycode: ''
            },
            rules2: {
                email: [{required: true, message: 'please enter your account', trigger: 'blur'}],
                password: [{required: true, message: 'enter your password', trigger: 'blur'}],
                verifycode: [{required: true, message: 'enter your verify code', trigger: 'blur'}]
            },
            checked: false
        }
    },
    methods: {
        handleSubmit(event){
            console.log(this.$refs)
            this.$refs.ruleForm2.validate((valid) => {
                if(valid){
                    this.logining = true;
                    if(this.ruleForm2.email === 'admin@gmail.com' && 
                       this.ruleForm2.password === '123456'){
                           this.logining = false;
                           sessionStorage.setItem('user', this.ruleForm2.email);
                           this.$router.push({path: '/reg'});
                           location.href = "/workbench.html"
                    }else{
                        this.logining = false;
                        this.$alert('email or password wrong!', 'info', {
                            confirmButtonText: 'ok'
                        })
                    }
                }else{
                    console.log('error submit!');
                    return false;
                }
            })
        }
    }
};
</script>

<style scoped>

.login-container {
    width: 100%;
    height: 100%;
}
.ghd-title {
    margin-bottom: 40px;
    text-align: center;
    font-size: 24px;
}
.login-page {
    -webkit-border-radius: 5px;
    border-radius: 5px;
    margin: 50px auto;
    width: 450px;
    padding: 55px 55px 15px;
    background: #fff;
    border: 1px solid #eaeaea;
    box-shadow: 0 0 25px #efebeb;
}
label.el-checkbox.rememberme {
    margin: 0px 0px 15px;
    text-align: left;
}

.el-link-btn {
    text-decoration: none;
    color: #DA5FB1;
}

a:hover {
    color: #DA5FB1 !important;
}

</style>