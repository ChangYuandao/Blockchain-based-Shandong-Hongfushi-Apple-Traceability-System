<template>
  <el-card>
    <h2>苹果物流运输信息录入</h2>
    <el-form :model="form" label-width="120px">
        <el-form-item label="山东红富士苹果产品批次">
          <el-input v-model="form.batchId" placeholder="如RFS-20250405-001" />
        </el-form-item>
      <el-form-item label="运输公司">
        <el-input
          v-model="form.shippingCompany"
          placeholder="请输入运输公司名称"
        />
      </el-form-item>

      <el-form-item label="运输单号">
        <el-input v-model="form.trackingNumber" placeholder="请输入运输单号" />
      </el-form-item>

      <el-form-item label="运输状态">
        <el-select v-model="form.shippingStatus" placeholder="请选择运输状态">
          <el-option label="待发货" value="待发货" />
          <el-option label="运输中" value="运输中" />
          <el-option label="已到达" value="已到达" />
          <el-option label="已完成" value="已完成" />
        </el-select>
      </el-form-item>

      <el-form-item label="预计到达日期">
        <el-date-picker
          v-model="form.estimatedArrivalDate"
          type="date"
          placeholder="请选择预计到达日期"
        />
      </el-form-item>

      <el-form-item label="实际到达日期">
        <el-date-picker
          v-model="form.actualArrivalDate"
          type="date"
          placeholder="请选择实际到达日期"
        />
      </el-form-item>

      <el-form-item label="运输方式">
        <el-select v-model="form.shippingMethod" placeholder="请选择运输方式">
          <el-option label="空运" value="空运" />
          <el-option label="海运" value="海运" />
          <el-option label="陆运" value="陆运" />
        </el-select>
      </el-form-item>

      <el-form-item label="运输费用">
        <el-input v-model="form.shippingCost" placeholder="请输入运输费用" />
      </el-form-item>


      <el-form-item>
        <el-button type="primary" @click="submit">提交</el-button>
        <el-button @click="reset">重置</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script>
import axios from "axios";
import { ElMessage, ElMessageBox } from "element-plus";

export default {
  data() {
    return {
      form: {
        batchId: "",
        shippingCompany: "",
        trackingNumber: "",
        shippingStatus: "",
        estimatedArrivalDate: "",
        actualArrivalDate: "",
        shippingMethod: "",
        shippingCost: "",
      },
    };
  },
  methods: {
    formatDate(date) {
      if (!date) return "";
      const d = new Date(date);
      return `${d.getFullYear()}-${(d.getMonth() + 1)
        .toString()
        .padStart(2, "0")}-${d.getDate().toString().padStart(2, "0")}`;
    },

    async submit() {
      try {
        await ElMessageBox.confirm("确认要提交物流信息吗？", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning",
        });

        // 格式化日期字段
        const payload = {
          ...this.form,
          estimatedArrivalDate: this.formatDate(this.form.estimatedArrivalDate),
          actualArrivalDate: this.formatDate(this.form.actualArrivalDate),
        };

        const res = await axios.post("/api/logistics", payload);

        if (res.data && res.data.message === "物流信息创建成功") {
          ElMessage.success("提交成功！");
          this.reset();
        } else {
          ElMessage.error(res.data.error || "提交失败！");
        }
      } catch (err) {
        if (err !== "cancel") {
          ElMessage.error("提交异常！");
          console.error(err);
        }
      }
    },

    reset() {
      this.form = {
        productName: "",
        shippingCompany: "",
        trackingNumber: "",
        shippingStatus: "",
        estimatedArrivalDate: "",
        actualArrivalDate: "",
        shippingMethod: "",
        shippingCost: "",
      };
    },
  },
};

</script>

<style scoped>
.el-card {
  height: 700;
  margin-top: 20px;
  padding: 20px;
}
</style>
