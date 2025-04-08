<template>
  <el-card :style="{ height: '500px' }">
    <h2>红富士苹果溯源信息录入</h2>
    <el-form :model="form" ref="form" label-width="120px">
      <el-form-item label="产品名称">
        <el-input v-model="form.productName" :disabled="true" />
      </el-form-item>

      <el-form-item label="产品批次编号">
        <el-input v-model="form.batchId" placeholder="如 RFS-20250405-001" />
      </el-form-item>

      <el-form-item label="种植基地名称">
        <el-input v-model="form.baseName" />
      </el-form-item>

      <el-form-item label="地理位置">
        <el-input
          v-model="form.location"
          placeholder="经纬度，如 37.5°N, 121.4°E"
        />
      </el-form-item>

      <el-form-item label="种植日期">
        <el-date-picker
          v-model="form.plantDate"
          type="date"
          placeholder="选择日期"
        />
      </el-form-item>

      <el-form-item label="预计采摘时间">
        <el-date-picker
          v-model="form.harvestDate"
          type="date"
          placeholder="选择日期"
        />
      </el-form-item>

      <el-form-item label="土壤类型">
        <el-select v-model="form.soilType" placeholder="请选择土壤类型">
          <el-option label="壤土" value="壤土" />
          <el-option label="沙土" value="沙土" />
          <el-option label="粘土" value="粘土" />
        </el-select>
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
        productName: "山东红富士苹果",
        batchId: "",
        baseName: "",
        location: "",
        plantDate: "",
        harvestDate: "",
        soilType: "",
      },
    };
  },
  methods: {
    formatDate(date) {
      if (!date) return "";
      const d = new Date(date);
      const year = d.getFullYear();
      const month = String(d.getMonth() + 1).padStart(2, "0");
      const day = String(d.getDate()).padStart(2, "0");
      return `${year}-${month}-${day}`; // 格式如 2025-04-16
    },
    async submit() {
      try {
        await ElMessageBox.confirm("确认要提交该产品信息吗？", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning",
        });

        const payload = {
          ...this.form,
          plantDate: this.formatDate(this.form.plantDate),
          harvestDate: this.formatDate(this.form.harvestDate),
        };

        const res = await axios.post(
          "/api/product",
          payload
        );

        console.log(res)
        if (res.data && res.data.message === "创建成功") {
          ElMessage.success("提交成功");
          this.reset();
        } else {
          ElMessage.error(res.data.error || "提交失败");
        }
      } catch (err) {
        if (err !== "cancel") {
          ElMessage.error("提交异常");
          console.error(err);
        }
      }
    },
    reset() {
      this.$refs.form.resetFields();
    },
  },
};
</script>
