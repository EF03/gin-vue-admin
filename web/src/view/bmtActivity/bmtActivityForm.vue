<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="活动主题:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="活动开始时间:">
          <el-date-picker v-model="formData.startTime" type="date" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="活动结束时间:">
          <el-date-picker v-model="formData.endTime" type="date" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="活动内容:">
          <el-input v-model="formData.content" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="活动花费:">
          <el-input v-model.number="formData.cost" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="活动人数:">
          <el-input v-model.number="formData.personNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BmtActivity'
}
</script>

<script setup>
import {
  createBmtActivity,
  updateBmtActivity,
  findBmtActivity
} from '@/api/bmtActivity'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
  name: '',
  startTime: new Date(),
  endTime: new Date(),
  content: '',
  cost: 0,
  personNum: 0,
})

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findBmtActivity({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.rebmtActivity
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createBmtActivity(formData.value)
      break
    case 'update':
      res = await updateBmtActivity(formData.value)
      break
    default:
      res = await createBmtActivity(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
  }
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
