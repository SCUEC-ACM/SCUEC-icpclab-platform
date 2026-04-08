<template>
  <div class="form-container">
    <h2>信息录入练习</h2>
    
    <div class="input-group">
      <label>姓名：</label>
      <input v-model="formData.name" type="text" placeholder="请输入姓名" />
    </div>

    <div class="input-group">
      <label>学号：</label>
      <input v-model="formData.student_id" type="text" placeholder="请输入学号" />
    </div>

    <div class="input-group">
      <label>联系方式：</label>
      <input v-model="formData.contact" type="text" placeholder="请输入手机或邮箱" />
    </div>

    <button @click="submitData">提交信息</button>
  </div>
</template>

<script setup>
defineOptions({name:'test'})
import { ref } from 'vue'

// 响应式变量，用于绑定输入框的数据
const formData = ref({
  name: '',
  student_id: '',
  contact: ''
})

// 提交函数
const submitData = async () => {
  // 简单的校验
  if (!formData.value.name || !formData.value.student_id) {
    alert('姓名和学号不能为空！')
    return
  }

  try {
    // 发送 POST 请求给 Go 后端
    const response = await fetch('http://localhost:3000/api/submit', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(formData.value)
    })

    const result = await response.json()

    if (response.ok) {
      alert(result.message) // 弹出成功提示
      // 清空输入框
      formData.value = { name: '', student_id: '', contact: '' }
    } else {
      alert(result.error) // 弹出错误提示（比如学号重复）
    }
  } catch (error) {
    console.error('请求失败:', error)
    alert('网络错误，请检查后端服务是否启动')
  }
}
</script>

<style scoped>
/* 一些简单的美化样式 */
.form-container { display: flex; flex-direction: column; width: 300px; gap: 10px; margin: 50px auto; }
.input-group { display: flex; justify-content: space-between; }
button { padding: 8px; cursor: pointer; background-color: #4CAF50; color: white; border: none; }
</style>