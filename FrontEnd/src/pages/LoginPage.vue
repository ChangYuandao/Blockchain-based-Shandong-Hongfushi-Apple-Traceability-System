<template>
  <div class="login-container">
    <div class="login-box">
      <h2>山东红富士苹果溯源系统</h2>

      <el-button type="primary" @click="loginWithMetaMask">
        使用 MetaMask 登录
      </el-button>

      <div style="margin: 10px 0; color: #999">或</div>

      <el-button type="success" @click="showRegister = true">
        注册新用户
      </el-button>
    </div>

    <!-- 注册弹窗 -->
    <el-dialog v-model="showRegister" title="注册新用户" width="30%">
      <div v-if="registerAddress">
        <p>
          钱包地址：<strong>{{ registerAddress }}</strong>
        </p>
        <el-select
          v-model="selectedRole"
          placeholder="请选择角色"
          style="width: 100%"
        >
          <el-option label="种植园" value="种植园" />
          <el-option label="物流企业" value="物流企业" />
          <el-option label="消费者" value="消费者" />
          <el-option label="管理员" value="管理员" />
        </el-select>
      </div>
      <template #footer>
        <el-button @click="showRegister = false">取消</el-button>
        <el-button type="primary" @click="registerUser">确认注册</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import axios from "axios";
import { ElMessage } from "element-plus";
import Web3 from "web3";

export default {
  name: "LoginPage",
  data() {
    return {
      registerAddress: "",
      selectedRole: "",
      showRegister: false,
      contractABI: [], // 存储合约的 ABI
      contractAddress: "", // 存储合约地址
    };
  },
  methods: {
    async loginWithMetaMask() {
      if (window.ethereum) {
        try {
          const accounts = await window.ethereum.request({
            method: "eth_requestAccounts",
          });
          const address = accounts[0];

          // 获取合约的 ABI 和地址
          await this.fetchContractDetails();

          const res = await axios.post("/api/login", { address });
          const role = res.data.data.role;

          console.log(res);
          ElMessage.success(`登录成功，角色：${role}`);

          // 保存角色到本地存储
          localStorage.setItem("role", role);

          this.redirectByRole(role); // 跳转到对应页面
        } catch (err) {
          ElMessage.error("登录失败");
        }
      } else {
        ElMessage.warning("请先安装 MetaMask 插件");
      }
    },
    async fetchContractDetails() {
      try {
        const response = await axios.get("/api/getContractDetails"); // 后端提供 ABI 和合约地址
        this.contractABI = response.data.abi;
        this.contractAddress = response.data.address;
      } catch (error) {
        ElMessage.error("获取合约信息失败");
      }
    },
    redirectByRole(role) {
      this.$nextTick(() => {
        switch (role) {
          case "种植园":
            this.$router.push("/submit"); // 种植园进入基础信息录入页面
            break;
          case "物流企业":
            this.$router.push("/process"); // 加工/物流企业进入补充过程数据页面
            break;
          case "管理员":
            this.$router.push("/audit"); // 管理员进入数据审核页面
            break;
          case "消费者":
            this.$router.push("/monitor"); // 消费者进入溯源查询页面
            break;
          default:
            this.$router.push("/login"); // 默认跳转到登录页面
            break;
        }
      });
    },
    async openRegisterDialog() {
      if (window.ethereum) {
        const accounts = await window.ethereum.request({
          method: "eth_requestAccounts",
        });
        this.registerAddress = accounts[0];
        this.showRegister = true;
      } else {
        ElMessage.warning("请先安装 MetaMask 插件");
      }
    },
    async registerUser() {
      if (this.selectedRole) {
        try {
          // 请求用户授权，获取当前账户
          const accounts = await window.ethereum.request({
            method: "eth_requestAccounts",
          });
          const currentAddress = accounts[0];

          // 发送到后端保存
          const res = await axios.post("/api/register", {
            address: currentAddress,
            role: this.selectedRole,
          });

          ElMessage.success("注册成功");
          this.showRegister = false;
        } catch (error) {
          console.error("注册失败", error);
          const message =
            error?.response?.data?.message || error?.message || "未知错误";
          ElMessage.error("注册失败：" + message);
        }
      } else {
        ElMessage.warning("请选择角色");
      }
    },
  },
  watch: {
    showRegister(val) {
      if (val) this.openRegisterDialog();
    },
  },
};
</script>

<style scoped>
.login-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background: linear-gradient(to right, #e0f7fa, #f1f8e9);
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
}

.login-box {
  background-color: white;
  padding: 40px;
  border-radius: 20px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  text-align: center;
  width: 360px;
}

.login-box h2 {
  margin-bottom: 30px;
  font-size: 24px;
  color: #2c3e50;
}

.el-button {
  width: 100%;
  margin-bottom: 15px;
  font-weight: bold;
}
</style>
