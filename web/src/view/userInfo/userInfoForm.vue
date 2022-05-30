<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="学号:">
          <el-input v-model="formData.xh" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:">
          <el-input v-model="formData.xm" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="性别:">
          <el-input v-model="formData.xb" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="民族:">
          <el-input v-model="formData.mz" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="生源地:">
          <el-input v-model="formData.syd" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出生日期:">
          <el-input v-model="formData.csrq" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份证号:">
          <el-input v-model="formData.sfzh" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="年级:">
          <el-input v-model="formData.nj" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="班级:">
          <el-input v-model="formData.bj" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="专业:">
          <el-input v-model="formData.zy" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="学院:">
          <el-input v-model="formData.xy" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="统一识别码:">
          <el-input v-model="formData.tysbm" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="专业号:">
          <el-input v-model="formData.zyh" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="院系号:">
          <el-input v-model="formData.yxh" clearable placeholder="请输入" />
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
  name: 'UserInfo'
}
</script>

<script setup>
import {
  createUserInfo,
  updateUserInfo,
  findUserInfo
} from '@/api/userInfo'

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
  xh: '',
  xm: '',
  xb: '',
  mz: '',
  syd: '',
  csrq: '',
  sfzh: '',
  nj: '',
  bj: '',
  zy: '',
  xy: '',
  tysbm: '',
  zyh: '',
  yxh: '',
  xjzt: '',
  pjjdcj: '',
})

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findUserInfo({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reuserInfo
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
      res = await createUserInfo(formData.value)
      break
    case 'update':
      res = await updateUserInfo(formData.value)
      break
    default:
      res = await createUserInfo(formData.value)
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
