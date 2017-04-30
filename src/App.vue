<template>
  <div id="app">
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span style="line-height: 36px;">动态更新</span>
      </div>
      <div class="service-selector">
        <el-select v-model="selectData.value" @change="selectChanged">
          <el-option
            v-for="item in selectData.options"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
        <el-button type="primary" :plain="true" :disabled="!isSelected"
         @click.native="addUpdateService">添加</el-button>
      </div>
      <div class="service-display-table">
        <el-table
          :data="tableData4"
          border
          style="width: 100%"
          max-height="250">
          <el-table-column
            prop="serviceName"
            label="服务名"
            width="150">
          </el-table-column>
          <el-table-column
            prop="imageName"
            label="镜像名">
          </el-table-column>
          <el-table-column
            prop="imageTag"
            label="镜像版本">
          </el-table-column>
          <el-table-column
            label="操作"
            width="120">
            <template scope="scope">
              <el-button
                @click.native.prevent="deleteRow(scope.$index, tableData4)"
                type="text"
                size="small">
                移除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div class="service-footer">
        <el-button type="primary" class="pull-right" @click.native="startHacking">更新</el-button>
        <div class="clearfix"></div>
      </div>
    </el-card>
    <!-- <el-button @click.native="startHacking">Let's do it</el-button> -->
  </div>
</template>

<script>

export default {
  data() {
    return {
      selectData: {
        options: [
          {
            value: '1',
            label: 'haha',
          },
          {
            value: '2',
            label: 'hehe',
          },
        ],
        value: '',
      },
      isSelected: false,
    };
  },
  methods: {
    selectChanged: function selectChanged() {
      if (this.selectData.value) {
        this.isSelected = true;
      }
    },
    addUpdateService: function addUpdateService() {
      this.isSelected = false;
      this.selectData.value = null;
    },
  },
};
  // methods: {
  //   startHacking () {
  //     this.$notify({
  //       title: 'It Works',
  //       message: 'We have laid the groundwork for you. Now it\'s your time to build something epic!',
  //       duration: 6000
  //     })
  //   }
  // }
</script>

<style>
body {
  font-family: Helvetica, sans-serif;
  -webkit-font-smoothing: antialiased;
}

.clearfix:before,
.clearfix:after {
    display: table;
    content: "";
}
.clearfix:after {
    clear: both
}

.el-card__header {
  background-color: #f5f7fa;
  font-size: 18px;
  font-weight: 400;
  color: #595f69;
}

.service-selector {
  margin-bottom: 10px
}

.service-display-table {
  margin-bottom: 10px
}

.el-table th>.cell {
  font-weight: 400;
  color: #9ba3af;
}

.pull-right {
  float: right;
}
</style>
