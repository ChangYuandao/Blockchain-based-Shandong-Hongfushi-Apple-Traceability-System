<template>
  <div class="p-6 bg-gray-50 rounded-lg shadow-md">
    <el-input
      v-model="batchId"
      placeholder="请输入批次编号（如：山东红富士苹果批次）"
      class="mb-4 w-full md:w-1/2"
      clearable
    >
      <template #append>
        <el-button type="primary" @click="fetchTraceInfo">查询</el-button>
      </template>
    </el-input>

    <!-- 产品信息表格 -->
    <div v-if="productData.length" class="mb-6">
      <h2 class="text-2xl font-semibold text-gray-800 mb-4">产品信息</h2>
      <el-table :data="productData" class="table-responsive">
        <el-table-column label="产品名称" prop="productName"></el-table-column>
        <el-table-column label="批次编号" prop="batchId"></el-table-column>
        <el-table-column label="基地名称" prop="baseName"></el-table-column>
        <el-table-column label="地理位置" prop="location"></el-table-column>
        <el-table-column label="种植日期" prop="plantDate"></el-table-column>
        <el-table-column label="收获日期" prop="harvestDate"></el-table-column>
        <el-table-column label="土壤类型" prop="soilType"></el-table-column>
        <el-table-column label="上链哈希" prop="TxHash"> </el-table-column>
      </el-table>
    </div>

    <!-- 物流信息表格 -->
    <div v-if="logisticsData.length">
      <h2 class="text-2xl font-semibold text-gray-800 mb-4">物流信息</h2>
      <el-table :data="logisticsData" class="table-responsive">
        <el-table-column label="运输公司" prop="shippingCompany"></el-table-column>
        <el-table-column label="批次编号" prop="batchId"></el-table-column>
        <el-table-column label="物流单号" prop="trackingNumber"></el-table-column>
        <el-table-column label="运输方式" prop="shippingMethod"></el-table-column>
        <el-table-column label="运费" prop="shippingCost"></el-table-column>
        <el-table-column label="预计到达" prop="estimatedArrivalDate"></el-table-column>
        <el-table-column label="实际到达" prop="actualArrivalDate"></el-table-column>
        <el-table-column label="运输状态" prop="shippingStatus"></el-table-column>
        <el-table-column label="上链哈希" prop="TxHash"></el-table-column>
      </el-table>
    </div>

    <el-empty
      v-if="!productData.length && !logisticsData.length && searched"
      description="暂无溯源信息"
    />
  </div>
</template>


<script>
import axios from "axios";
export default {
  data() {
    return {
      batchId: "",
      productData: [],
      logisticsData: [],
      searched: false,
    };
  },
  methods: {
    // 查询产品和物流信息
    async fetchTraceInfo() {
      if (!this.batchId) return;

      try {
        const [productRes, logisticsRes] = await Promise.all([
          axios.post(`/api/GetProductByBatchId`, { batchId: this.batchId }), // 使用 POST 请求
          axios.post(`/api/GetLogisticsByBatchId`, { batchId: this.batchId }), // 使用 POST 请求
        ]);
        this.productData = productRes.data?.data || [];
        this.logisticsData = logisticsRes.data?.data || [];
        // 打印返回的数据，调试检查
        console.log("Product Data:", productRes.data?.data);
        console.log("Logistics Data:", logisticsRes.data?.data);
      } catch (error) {
        console.error("查询失败", error);
      } finally {
        this.searched = true;
      }
    },

    // 获取状态标签类型
    getStatusType(status) {
      switch (status) {
        case "approved":
          return "success";
        case "pending":
          return "warning";
        case "rejected":
          return "danger";
        default:
          return "info";
      }
    },
  },
};
</script>

<style scoped>
.table-responsive {
  width: 100%;
  overflow-x: auto;
}

.el-input {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
}

.el-button {
  padding: 8px 16px;
  border-radius: 4px;
  background-color: #409EFF;
  color: white;
}

.el-button:hover {
  background-color: #66b1ff;
}

h2 {
  font-size: 1.5rem;
  font-weight: 600;
  color: #2d3748;
}

.el-tag {
  font-size: 0.875rem;
}

.el-link {
  color: #409EFF;
  text-decoration: underline;
}

.el-link:hover {
  color: #1e88e5;
}

.el-empty {
  text-align: center;
}

@media (max-width: 768px) {
  .el-input {
    width: 100%;
  }
}
</style>
