<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="羽球活动主键:">
          <el-input v-model.number="formData.bmtId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="红对队员1:">
          <el-select v-model="formData.redUserId1" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in bmtUlOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="红对队员2:">
          <el-select v-model="formData.redUserId2" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in bmtUlOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="蓝对队员1:">
          <el-select v-model="formData.blueUserId1" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in bmtUlOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="蓝对队员2:">
          <el-select v-model="formData.blueUserId2" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in bmtUlOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="红对分数:">
          <el-input v-model.number="formData.redScore" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="藍隊分數:">
          <el-input v-model.number="formData.blueScore" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="胜负:">
          <el-select v-model="formData.winner" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in winnerOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'BmtPeriod'
}
</script>

<script setup>
import {
  createBmtPeriod,
  updateBmtPeriod,
  findBmtPeriod
} from '@/api/bmtPeriod'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const bmtUlOptions = ref([])
const winnerOptions = ref([])
const formData = ref({
        bmtId: 0,
        redUserId1: undefined,
        redUserId2: undefined,
        blueUserId1: undefined,
        blueUserId2: undefined,
        redScore: 0,
        blueScore: 0,
        winner: undefined,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBmtPeriod({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rebmtPeriod
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    bmtUlOptions.value = await getDictFunc('bmtUl')
    winnerOptions.value = await getDictFunc('winner')
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createBmtPeriod(formData.value)
          break
        case 'update':
          res = await updateBmtPeriod(formData.value)
          break
        default:
          res = await createBmtPeriod(formData.value)
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
