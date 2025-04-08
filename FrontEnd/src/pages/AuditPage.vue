<template>
  <div>
    <el-tabs v-model="activeTab">
      <el-tab-pane label="种植园数据" name="plant">
        <el-table :data="plantData" stripe height="500px">
          <el-table-column label="产品名称" prop="productName" />
          <el-table-column label="批次号" prop="batchId" />
          <el-table-column label="基地名称" prop="baseName" />
          <el-table-column label="所在地" prop="location" />
          <el-table-column label="种植日期" prop="plantDate" />
          <el-table-column label="收获日期" prop="harvestDate" />
          <el-table-column label="土壤类型" prop="soilType" />
          <el-table-column label="审核状态" :formatter="formatStatus" />
          <el-table-column label="操作">
            <template #default="scope">
              <el-button
                @click="approveData(scope.row, 'plant')"
                size="small"
                type="success"
                >通过</el-button
              >
              <el-button
                @click="rejectData(scope.row, 'plant')"
                size="small"
                type="danger"
                >拒绝</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
      <el-tab-pane label="物流数据" name="shipping">
        <el-table :data="shippingData" stripe height="500px">
          <el-table-column label="山东红富士苹果批次" prop="batchId" width="200px"/>
          <el-table-column label="物流公司" prop="shippingCompany" />
          <el-table-column label="追踪号" prop="trackingNumber" />
          <el-table-column label="运输状态" prop="shippingStatus" />
          <el-table-column label="预计到达时间" prop="estimatedArrivalDate" />
          <el-table-column label="实际到达时间" prop="actualArrivalDate" />
          <el-table-column label="运输方式" prop="shippingMethod" />
          <el-table-column label="运输费用" prop="shippingCost" />
          <el-table-column label="审核状态" :formatter="formatStatus" />
          <el-table-column label="操作" width="200px">
            <template #default="scope">
              <el-button
                @click="approveData(scope.row, 'shipping')"
                size="small"
                type="success"
                >通过</el-button
              >
              <el-button
                @click="rejectData(scope.row, 'shipping')"
                size="small"
                type="danger"
                >拒绝</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import {
  ElMessage,
  ElButton,
  ElTable,
  ElTableColumn,
  ElTabs,
  ElTabPane,
} from "element-plus";
import Web3 from "web3";
import axios from "axios";
import ABI from "../assets/abis.json";

export default {
  components: { ElButton, ElTable, ElTableColumn, ElTabs, ElTabPane },
  data() {
    return {
      activeTab: "plant",
      plantData: [],
      shippingData: [],
    };
  },
  methods: {
    // 格式化审核状态
    formatStatus(row) {
      return row.status === "pending"
        ? "待审核"
        : row.status === "approved"
        ? "已通过"
        : "已拒绝";
    },

    // 获取数据
    async fetchData() {
      try {
        const [plantRes, shippingRes] = await Promise.all([
          axios.get("/api/getAllProducts"),
          axios.get("/api/getAllLogistics"),
        ]);
        this.plantData = plantRes.data.data || [];
        this.shippingData = shippingRes.data.data || [];
      } catch (error) {
        console.error("获取数据失败", error);
      }
    },
    async approveData(row, type) {
      const web3 = new Web3(window.ethereum);
      await window.ethereum.request({ method: "eth_requestAccounts" });

      const contract = new web3.eth.Contract(
        ABI,
        "0x6Aa6ADAA650Db64e4e2b7cCFB8737f94364c3F43"
      );

      const accounts = await web3.eth.getAccounts();
      const fromAddress = accounts[0];

      try {
        let batchId = row.batch_id || row.batchId || ""; // 兼容不同字段命名
        let role = type === "plant" ? "plant" : "shipping"; // 修改角色为 'plant' 或 'shipping'
        let data = JSON.stringify(row); // 将整条数据序列化上传

        const tx = await contract.methods
          .uploadSegment(batchId, role, data)
          .send({ from: fromAddress });

        console.log(
          `${type === "plant" ? "种植园" : "物流"}数据审核通过，交易ID：`,
          tx.transactionHash
        );

        // 更新行数据
        row.txHash = tx.transactionHash;
        row.status = "approved";

        const url =
          type === "plant"
            ? `/api/updateProductStatus`
            : `/api/updateShippingStatus`;

        await axios.post(url, {
          id: row.id,
          status: "approved",
          txHash: tx.transactionHash,
        });

        ElMessage({
          message: "数据已成功上链",
          type: "success",
        });
        this.fetchData(); // 刷新数据
      } catch (error) {
        console.error("上链失败", error);
        ElMessage({
          message: "数据已成功上链",
          type: "error  ",
        });
      }
    },
    // 审核拒绝
    async rejectData(row, type) {
      // 处理拒绝逻辑，比如更新数据库状态等
      await fetch(`/api/update-status?id=${row.id}&status=rejected`, {
        method: "POST",
      });

      row.status = "rejected";
      alert(`${type === "plant" ? "种植园" : "物流"}数据已拒绝`);
      this.fetchData(); // 更新数据
    },
  },
  mounted() {
    this.fetchData();
  },
};
</script>

<style scoped>
/* 页面外层容器 */
.el-tabs {
  margin-top: 20px;
  padding: 20px;
  border-radius: 10px;
  background-color: #f9f9f9;
}

/* 选项卡容器样式 */
.el-tab-pane {
  padding: 10px;
}

/* 表格样式 */
.el-table {
  margin-top: 20px;
  border-radius: 10px;
  border: 1px solid #ddd;
}

.el-table th {
  background-color: #f1f1f1;
  color: #333;
}

.el-table .cell {
  word-break: break-word;
  word-wrap: break-word;
}

/* 操作按钮样式 */
.el-button {
  margin-right: 10px;
}

.el-button.success {
  background-color: #67c23a;
  border-color: #67c23a;
}

.el-button.danger {
  background-color: #f56c6c;
  border-color: #f56c6c;
}

/* 表格列内容 */
.el-table-column {
  text-align: center;
}

.el-table-column .el-button {
  margin-left: 10px;
}

/* 审核状态 */
.el-tag {
  margin-top: 5px;
}

/* 页面标题 */
h2 {
  font-size: 24px;
  color: #333;
  margin-bottom: 20px;
}

/* 页面底部 */
footer {
  margin-top: 30px;
  text-align: center;
  padding: 20px;
  background-color: #f1f1f1;
  border-radius: 10px;
}

/* 弹出框样式 */
.el-dialog {
  border-radius: 10px;
}

/* 修改按钮 hover 效果 */
.el-button:hover {
  opacity: 0.8;
}

/* 自定义表格列宽 */
.el-table-column:nth-child(1),
.el-table-column:nth-child(2) {
  width: 150px;
}

.el-table-column:nth-child(3),
.el-table-column:nth-child(4) {
  width: 200px;
}

.el-table-column:nth-child(5),
.el-table-column:nth-child(6) {
  width: 180px;
}
</style>
