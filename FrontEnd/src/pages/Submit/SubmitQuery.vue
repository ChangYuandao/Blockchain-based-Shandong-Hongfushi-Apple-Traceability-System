<template>
    <el-card>
      <h2>溯源信息查询</h2>
      <el-form :inline="true" @submit.prevent>
        <el-form-item label="溯源码">
          <el-input v-model="traceCode" placeholder="输入溯源码" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="query">查询</el-button>
        </el-form-item>
      </el-form>
  
      <el-divider />
  
      <div v-if="info">
        <el-descriptions title="查询结果" :column="2" border>
          <el-descriptions-item label="产品名称">{{ info.productName }}</el-descriptions-item>
          <el-descriptions-item label="批次编号">{{ info.batchId }}</el-descriptions-item>
          <el-descriptions-item label="基地名称">{{ info.baseName }}</el-descriptions-item>
          <el-descriptions-item label="地理位置">{{ info.location }}</el-descriptions-item>
          <el-descriptions-item label="种植日期">{{ info.plantDate }}</el-descriptions-item>
          <el-descriptions-item label="采摘时间">{{ info.harvestDate }}</el-descriptions-item>
          <el-descriptions-item label="土壤类型">{{ info.soilType }}</el-descriptions-item>
          <el-descriptions-item label="农药记录">{{ info.pesticide }}</el-descriptions-item>
          <el-descriptions-item label="肥料记录">{{ info.fertilizer }}</el-descriptions-item>
          <el-descriptions-item label="备注">{{ info.remark }}</el-descriptions-item>
          <el-descriptions-item label="状态">{{ info.status }}</el-descriptions-item>
        </el-descriptions>
      </div>
    </el-card>
  </template>
  
  <script>
  import { ethers } from 'ethers'
  import { ElMessage } from 'element-plus'
  
  export default {
    data() {
      return {
        traceCode: '',  // 溯源码
        info: null,     // 查询结果
        provider: null, // 用于存储Ethers提供者
        contract: null, // 用于存储合约实例
      }
    },
    methods: {
      async query() {
        if (!this.traceCode) {
          ElMessage.error('请输入溯源码')
          return
        }
  
        // 连接到MetaMask
        if (typeof window.ethereum !== 'undefined') {
          try {
            // 获取用户地址
            await window.ethereum.request({ method: 'eth_requestAccounts' })
            this.provider = new ethers.JsonRpcProvider(window.ethereum)
            const signer = this.provider.getSigner()
  
            // 连接到你的合约，使用合约地址和ABI
            const contractAddress = '0xYourContractAddress'  // 你的合约地址
            const abi = [
              // 合约的ABI接口
              "function getTraceByCode(string memory traceCode) public view returns (string memory, string memory, string memory, string memory, string memory, string memory, string memory, string memory, string memory)"
            ]
            this.contract = new ethers.Contract(contractAddress, abi, signer)
  
            // 调用智能合约查询溯源信息
            const trace = await this.contract.getTraceByCode(this.traceCode)
  
            // 格式化返回的数据
            this.info = {
              productName: trace[0],
              batchId: trace[1],
              baseName: trace[2],
              location: trace[3],
              plantDate: trace[4],
              harvestDate: trace[5],
              soilType: trace[6],
              pesticide: trace[7],
              fertilizer: trace[8],
              remark: trace[9],
              status: "已上链", // 假设状态为已上链
            }
          } catch (err) {
            console.error(err)
            ElMessage.error('查询失败，请检查钱包连接或合约是否正确')
          }
        } else {
          ElMessage.error('MetaMask 未连接')
        }
      }
    }
  }
  </script>
  
  <style scoped>
  .el-card {
    max-width: 900px;
    margin: 20px auto;
  }
  </style>
  