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
  name: 'WeMrdk2'
}
</script>

<script setup>
import {
  createWeMrdk2,
  updateWeMrdk2,
  findWeMrdk2
} from '@/api/weMrdk2'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWeMrdk2({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reweMrdk2
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
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
