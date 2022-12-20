<template>
    <div>
        <el-row style="margin-bottom: 10px;">
            <el-button @click="(currentProxyApiType = getProxyApiTypes.all) && initProxyTableData()"
                :type="currentProxyApiType === getProxyApiTypes.all ? 'primary' : 'info'">全部代理</el-button>
            <el-button @click="(currentProxyApiType = getProxyApiTypes.old) && initProxyTableData()"
                       :type="currentProxyApiType === getProxyApiTypes.old ? 'primary' : 'info'">初始定义代理</el-button>
<!--            <el-button :type="currentProxyApiType === getProxyApiTypes.new ? 'primary' : 'info'">临时代理</el-button>-->
        </el-row>
        <el-table
            stripe
            border
            :data="tableData"
            style="width: 100%">
            <el-table-column type="expand">
                <template slot-scope="props">
                    <el-descriptions class="margin-top" size="mini" border style="margin: 0px 10px;">
                        <el-descriptions-item v-for="p in tableProps.description" :key="p.prop" :label="p.label">
                            <span v-if="!p.bool">{{ props.row[p.prop] }}</span>
                            <span v-else>
                                <el-tag v-if="props.row[p.prop]"
                                        effect="dark"
                                        type="success">是</el-tag>
                                <el-tag v-else effect="dark"
                                        type="danger">否</el-tag>
                            </span>
                        </el-descriptions-item>
                    </el-descriptions>
                </template>
            </el-table-column>
            <el-table-column
                v-for="p in tableProps.normal"
                :key="p.prop"
                :prop="p.prop"
                :label="p.label">
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import {
    getProxyApi,
    getProxyApiTypes,
} from "../../apis";

export default {
    name: "index",
    data() {
        return {
            getProxyApiTypes,
            currentProxyApiType: getProxyApiTypes.old,
            proxyApiParentId: '',
            tableProps: {
                normal: [
                    {
                        prop: 'id',
                        label: 'Id',
                    },
                    {
                        prop: 'Group',
                        label: '组名',
                    },
                    {
                        prop: 'Path',
                        label: '请求路径',
                    },
                    {
                        prop: 'ParentId',
                        label: '父级id',
                    },
                ],
                description: [
                    {
                        prop: 'Uri',
                        label: '代理链接',
                    },
                    {
                        prop: 'Method',
                        label: '请求方式',
                    },
                    {
                        prop: 'Cache',
                        label: '是否进行缓存',
                        bool: true,
                    },
                    {
                        prop: 'CacheByURI',
                        label: '添加请求参数作为缓存',
                        bool: true,
                    },
                    {
                        prop: 'CacheWithBody',
                        label: '添加请求体作为缓存',
                        bool: true,
                    },
                    {
                        prop: 'New',
                        label: '是否未添加',
                        bool: true,
                    }
                ]
            },
            tableData: [
                {
                    "id": '',
                    "New": "true",
                    "Group": "table",
                    "Path": "system/dict/data/type/sys_oper_type",
                    "Method": "GET",
                    "Uri": "http://172.20.109.155:8080/",
                    "Cache": "true",
                    "CacheByURI": "true",
                    "CacheWithBody": "false",
                    "ParentId": "#table_$_*_$_GET#"
                }
            ]
        }
    },
    methods: {
        initProxyTableData() {
            getProxyApi(this.currentProxyApiType,this.proxyApiParentId).then(ret => {
                let arr = [];
                for (let i in ret.Data) {
                    arr.push({
                        ...ret.Data[i],
                        id: i
                    })
                }
                // console.log(ret);
                this.tableData.splice(0,this.tableData.length,...arr);
            });
        }
    },
    mounted() {
        this.initProxyTableData();
    }
}
</script>

<style scoped>

</style>