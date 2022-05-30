<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用户日志id:">
          <el-input v-model.number="formData.logId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="学号:">
          <el-input v-model="formData.xh" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="目前居住地:">
          <el-input v-model="formData.szdq" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否与湖北人员接触:">
          <el-input v-model="formData.sfhxdgr" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="健康状况:">
          <el-input v-model="formData.stsfjk" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="今日体温:">
          <el-input v-model="formData.jrtw" clearable placeholder="请输入" />
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
  name: 'WeMrdk1'
}
</script>

<script setup>
import {
  createWeMrdk1,
  updateWeMrdk1,
  findWeMrdk1
} from '@/api/weMrdk1'

// 自动获取字典
// eslint-disable-next-line no-unused-vars
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
  logId: 0,
  xh: '',
  name: '',
  szdq: '',
  sfhxdgr: '',
  stsfjk: '',
  jrtw: '',
})

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findWeMrdk1({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reweMrdk1
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
      res = await createWeMrdk1(formData.value)
      break
    case 'update':
      res = await updateWeMrdk1(formData.value)
      break
    default:
      res = await createWeMrdk1(formData.value)
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
