<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="学号">
          <el-input v-model="searchInfo.xh" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="姓名">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
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
        <el-table-column align="left" label="学号" prop="xh" width="120" />
        <el-table-column align="left" label="姓名" prop="name" width="120" />
        <el-table-column align="left" label="性别" prop="xb" width="120" />
        <el-table-column align="center" label="目前居住地" prop="szdq" width="170" />
        <el-table-column align="left" label="是否与湖北人员接触" prop="sfhxdgr" width="160" />
        <el-table-column align="left" label="健康状况" prop="stsfjk" width="120" />
        <el-table-column align="left" label="今日体温" prop="jrtw" width="120" />
        <el-table-column align="left" label="是否发热" prop="sffr" width="120" />
        <el-table-column align="left" label="是否咳嗽" prop="sfks" width="120" />
        <el-table-column align="left" label="是否胸闷" prop="sfxm" width="120" />
        <el-table-column align="left" label="其他异常情况" prop="qtycqk" width="120" />
        <el-table-column align="center" label="14天内是否有中高风险地区旅居史" prop="ywjchblj" width="120" />
        <el-table-column align="center" label="目前居住地新冠肺炎疫情风险等级" prop="ywjcqzbl" width="120" />
        <el-table-column align="center" label="14天内是否接触中高风险地区旅居史人员" prop="xjzdywqzbl" width="120" />
        <el-table-column align="center" label="今日体温是否正常" prop="twsfzc" width="140" />
        <el-table-column align="center" label="今日是否有与新冠病毒感染有关的症状" prop="ywytdzz" width="120" />
        <el-table-column align="left" label="寒假是否离渝" prop="hjsfly" width="120" />
        <el-table-column align="left" label="是否已返渝" prop="sfyfy" width="120" />
        <el-table-column align="left" label="反渝路线（途径哪些地方）" prop="fylx" width="120" />
        <el-table-column align="left" label="返回乘坐的交通工具（班次车牌）" prop="fyjtgj" width="120" />
        <el-table-column align="left" label="返渝到达时间" prop="fyddsj" width="120" />
        <el-table-column align="left" label="是否报告社区" prop="sfbgsq" width="120" />
        <el-table-column align="left" label="是否居家隔离" prop="sfjjgl" width="120" />
        <el-table-column align="left" label="居家隔离起始时间" prop="jjglqssj" width="120" />
        <el-table-column align="left" label="未居家隔离的目前去向" prop="wjjglmqqx" width="120" />
        <el-table-column align="left" label="详细地址" prop="xxdz" width="120" />
        <el-table-column align="left" label="本人及家庭成员是否为确诊病例" prop="brsfqz" width="120" />
        <el-table-column align="left" label="本人及家庭成员是否为疑似病例" prop="brsfys" width="120" />
        <el-table-column align="left" label="有无疾病史" prop="jbs" width="120" />
        <el-table-column align="center" label="健康码识别信息" prop="jkmresult" width="120" />
        <el-table-column align="left" label="按钮组" width="120">
          <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateWeMrdk2Func(scope.row)">变更</el-button>
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
        <el-form-item label="用户日志id:">
          <el-input v-model.number="formData.logId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="学号:">
          <el-input v-model="formData.xh" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="性别:">
          <el-input v-model="formData.xb" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系电话:">
          <el-input v-model="formData.lxdh" clearable placeholder="请输入" />
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
        <el-form-item label="是否发热:">
          <el-input v-model="formData.sffr" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否咳嗽:">
          <el-input v-model="formData.sfks" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否胸闷:">
          <el-input v-model="formData.sfxm" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="其他异常情况:">
          <el-input v-model="formData.qtycqk" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="14天内是否有中高风险地区旅居史:">
          <el-input v-model="formData.ywjchblj" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="目前居住地新冠肺炎疫情风险等级:">
          <el-input v-model="formData.ywjcqzbl" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="14天内是否接触中高风险地区旅居史人员:">
          <el-input v-model="formData.xjzdywqzbl" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="今日体温是否正常:">
          <el-input v-model="formData.twsfzc" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="今日是否有与新冠病毒感染有关的症状:">
          <el-input v-model="formData.ywytdzz" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否咳嗽:">
          <el-input v-model="formData.jbsks" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否乏力:">
          <el-input v-model="formData.jbsfl" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否鼻塞:">
          <el-input v-model="formData.jbsbs" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否流涕:">
          <el-input v-model="formData.jbslt" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否咽痛:">
          <el-input v-model="formData.jbsyt" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否腹泻:">
          <el-input v-model="formData.jbsfx" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="寒假是否离渝:">
          <el-input v-model="formData.hjsfly" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否已返渝:">
          <el-input v-model="formData.sfyfy" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="反渝路线（途径哪些地方）:">
          <el-input v-model="formData.fylx" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="返回乘坐的交通工具（班次车牌）:">
          <el-input v-model="formData.fyjtgj" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="返渝到达时间:">
          <el-input v-model="formData.fyddsj" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否报告社区:">
          <el-input v-model="formData.sfbgsq" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否居家隔离:">
          <el-input v-model="formData.sfjjgl" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="居家隔离起始时间:">
          <el-input v-model="formData.jjglqssj" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="未居家隔离的目前去向:">
          <el-input v-model="formData.wjjglmqqx" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="详细地址:">
          <el-input v-model="formData.xxdz" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="本人及家庭成员是否为确诊病例:">
          <el-input v-model="formData.brsfqz" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="本人及家庭成员是否为疑似病例:">
          <el-input v-model="formData.brsfys" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="有无疾病史:">
          <el-input v-model="formData.jbs" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="健康码识别信息:">
          <el-input v-model="formData.jkmresult" clearable placeholder="请输入" />
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
  name: 'WeMrdk2'
}
</script>

<script setup>
import {
  createWeMrdk2,
  deleteWeMrdk2,
  deleteWeMrdk2ByIds,
  updateWeMrdk2,
  findWeMrdk2,
  getWeMrdk2List
} from '@/api/weMrdk2'

// 全量引入格式化工具 请按需保留
// eslint-disable-next-line no-unused-vars
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  logId: 0,
  xh: '',
  name: '',
  xb: '',
  lxdh: '',
  szdq: '',
  sfhxdgr: '',
  stsfjk: '',
  jrtw: '',
  sffr: '',
  sfks: '',
  sfxm: '',
  qtycqk: '',
  ywjchblj: '',
  ywjcqzbl: '',
  xjzdywqzbl: '',
  twsfzc: '',
  ywytdzz: '',
  jbsks: '',
  jbsfl: '',
  jbsbs: '',
  jbslt: '',
  jbsyt: '',
  jbsfx: '',
  hjsfly: '',
  sfyfy: '',
  fylx: '',
  fyjtgj: '',
  fyddsj: '',
  sfbgsq: '',
  sfjjgl: '',
  jjglqssj: '',
  wjjglmqqx: '',
  xxdz: '',
  brsfqz: '',
  brsfys: '',
  jbs: '',
  jkmresult: '',
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

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
  const table = await getWeMrdk2List({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteWeMrdk2Func(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const log_id = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          log_id.push(item.ID)
        })
  const res = await deleteWeMrdk2ByIds({ log_id })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === log_id.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateWeMrdk2Func = async(row) => {
  const res = await findWeMrdk2({ logId: row.logId })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reweMrdk2
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteWeMrdk2Func = async(row) => {
  const res = await deleteWeMrdk2({ logId: row.logId })
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
    logId: 0,
    xh: '',
    name: '',
    xb: '',
    lxdh: '',
    szdq: '',
    sfhxdgr: '',
    stsfjk: '',
    jrtw: '',
    sffr: '',
    sfks: '',
    sfxm: '',
    qtycqk: '',
    ywjchblj: '',
    ywjcqzbl: '',
    xjzdywqzbl: '',
    twsfzc: '',
    ywytdzz: '',
    jbsks: '',
    jbsfl: '',
    jbsbs: '',
    jbslt: '',
    jbsyt: '',
    jbsfx: '',
    hjsfly: '',
    sfyfy: '',
    fylx: '',
    fyjtgj: '',
    fyddsj: '',
    sfbgsq: '',
    sfjjgl: '',
    jjglqssj: '',
    wjjglmqqx: '',
    xxdz: '',
    brsfqz: '',
    brsfys: '',
    jbs: '',
    jkmresult: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createWeMrdk2(formData.value)
      break
    case 'update':
      res = await updateWeMrdk2(formData.value)
      break
    default:
      res = await createWeMrdk2(formData.value)
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
