<template>
  <div id="app">
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span style="line-height: 36px;">动态更新</span>
      </div>
      <div class="service-selector">
        <el-select v-model="selectedOption" @change="selectChanged">
          <el-option
            v-for="service in filteredAppServices"
            :key="service.ID"
            :label="service.Spec.Name"
            :value="service.ID">
          </el-option>
        </el-select>
        <el-button type="primary" :plain="true" :disabled="!isSelected"
         @click.native="addUpdateService">添加</el-button>
      </div>
      <div class="service-display-table">
        <el-table
          :data="updateServices"
          border
          style="width: 100%"
          max-height="250">
          <el-table-column
            prop="Spec.Name"
            label="服务名"
            width="200">
          </el-table-column>
          <el-table-column
            label="镜像">
            <template scope="scope">
              <el-input v-model="scope.row.Spec.TaskTemplate.ContainerSpec.Image" placeholder="请输入内容"></el-input>
            </template>
          </el-table-column>
          <el-table-column
            label="操作"
            width="120">
            <template scope="scope">
              <el-button
                @click.native.prevent="deleteUpdateService(scope.$index)"
                type="text"
                size="small">
                移除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div class="service-footer">
        <el-button type="primary" class="pull-right" @click.native="updateApp">更新</el-button>
        <div class="clearfix"></div>
      </div>
    </el-card>
  </div>
</template>

<script>
import * as _ from './helpers/lodash';

const queryString = require('query-string');
/* eslint-disable func-names */
export default {
  data() {
    return {
      appName: '',
      appServices: [
        {
          ID: '123',
          Spec: {
            Name: 'haha',
            TaskTemplate: {
              ContainerSpec: {
                Image: '1.1.1.1/aa/2038:0.1.1@12388',
              },
            },
          },
        },
        {
          ID: '222',
          Spec: {
            Name: 'heihei',
            TaskTemplate: {
              ContainerSpec: {
                Image: 'aa/2038:0.1.1@12388',
              },
            },
          },
        },
      ],
      updateServices: [],
      selectedOption: '',
      selectData: {
        options: [],
        value: '',
      },
      isSelected: false,
    };
  },
  computed: {
    // a computed getter
    mapedAppServices() {
      const maped = {};
      for (const i in this.appServices) {
        if (this.appServices.hasOwnProperty(i)) {
          maped[this.appServices[i].ID] = this.appServices[i];
        }
      }
      return maped;
    },
    filteredAppServices() {
      return _.filter(this.appServices, (s) => {
        for (const us of this.updateServices) {
          if (us.ID === s.ID) {
            return false;
          }
        }
        return true;
      });
    },
  },
  created() {
    this.appName = queryString.parse(location.search).ObjectId;
    this.bindAppServices(this.appName);
  },
  methods: {
    // 绑定数据源，初始化 App 的服务
    bindAppServices(appName) {
      // GET request
      this.$http.get(`api/app/${appName}`).then(response => {
        // get body data
        this.appServices = response.body;
      }, response => {
        console.log(response);
      });
    },

    // 当选择改变时产生的事件
    selectChanged() {
      if (this.selectedOption) {
        this.isSelected = true;
      }
    },

    // 点击添加按钮
    addUpdateService() {
      this.isSelected = false;
      const updateService = _.cloneDeep(this.mapedAppServices[this.selectedOption]);
      const updateServiceImage = updateService.Spec.TaskTemplate.ContainerSpec.Image;
      updateService.Spec.TaskTemplate.ContainerSpec.Image = updateServiceImage.slice(0, updateServiceImage.lastIndexOf('@'));
      this.updateServices.push(updateService);
      this.selectedOption = null;
      console.log(this.updateServices);
    },

    // 删除某一行
    deleteUpdateService(index) {
      const ss = this.updateServices[index];
      this.updateServices = _.reject(this.updateServices, (s) => {
        return s.ID === ss.ID;
      });
    },

    // 更新 App
    updateApp() {
      const vm = this;
      const jsonData = [];
      for (const us of this.updateServices) {
        jsonData.push({
          ID: us.ID,
          Image: us.Spec.TaskTemplate.ContainerSpec.Image,
        });
      }
      this.$http.put(`api/app/${this.appName}`, jsonData).then(response => {
        // get body data
        console.log('error:', response);
        vm.$message('应用开始更新，点击服务查看详情');
      }, response => {
        vm.$message('应用开始更新，点击服务查看详情');
        console.log('error:', response);
      });
    },
  },
};
</script>

<style>
body {
  font-family: Helvetica, sans-serif;
  -webkit-font-smoothing: antialiased;
  margin: 0px;
}

.clearfix:before,
.clearfix:after {
    display: table;
    content: "";
}
.clearfix:after {
    clear: both
}

.el-card {
  box-shadow: 0 1px 4px rgba(204,209,217,.3);
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
