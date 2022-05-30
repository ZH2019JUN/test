<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="姓名">
          <el-input v-model="searchInfo.xm" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="班级">
          <el-input v-model="searchInfo.bj" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="专业">
          <el-input v-model="searchInfo.zy" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
            <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="mini" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="学号" prop="xh" width="100" />
        <el-table-column align="left" label="姓名" prop="xm" width="80" />
        <el-table-column align="left" label="民族" prop="mz" width="80" />
        <el-table-column align="left" label="生源地" prop="syd" width="80" />
        <el-table-column align="center" label="身份证号" prop="sfzh" width="170" />
        <el-table-column align="center" label="年级" prop="nj" width="80" />
        <el-table-column align="center" label="班级" prop="bj" width="100" />
        <el-table-column align="center" label="专业" prop="zy" width="170" />
        <el-table-column align="center" label="按钮组" width="170">
          <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateUserInfoFunc(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
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
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
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
  deleteUserInfo,
  deleteUserInfoByIds,
  updateUserInfo,
  findUserInfo,
  getUserInfoList
} from '@/api/userInfo'

// 全量引入格式化工具 请按需保留
// eslint-disable-next-line no-unused-vars
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
// eslint-disable-next-line no-unused-vars
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

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const xy = ref('体育学院')

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getUserInfoList({ page: page.value, pageSize: pageSize.value, xy: xy.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteUserInfoFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
  multipleSelection.value.map(item => {
    ids.push(item.ID)
  })
  const res = await deleteUserInfoByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateUserInfoFunc = async(row) => {
  const res = await findUserInfo({ xh: row.xh })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reuserInfo
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteUserInfoFunc = async(row) => {
  const res = await deleteUserInfo({ xh: row.xh })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
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
  }
}
// 弹窗确定
const enterDialog = async() => {
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
    closeDialog()
    getTableData()
  }
}
</script>

<style>
</style>
