<template>
    <el-container style="width:100vw;height: 100vh;background: #f4f4f4;">
        <el-header style="background: #1e88e5;"><span class="app-header">代理管理</span>
        </el-header>
        <el-container style="height: calc(100vh - 60px);">
            <el-aside width="65px" style="border-right: 1px solid gray;">
                <div class="app-el-menu">
                    <el-tooltip class="app-el-menu-li" effect="dark" content="代理"
                                placement="right">
<!--                        <span @click="">-->
                        <span>
                            <i class="el-icon-s-promotion"></i>
                        </span>
                    </el-tooltip>
                </div>
            </el-aside>
            <el-container>
                <el-main>
<!--                    <component :is="current"></component>-->
                    <ProxyApi></ProxyApi>
                    <div class="app-main">
                    </div>
                </el-main>
            </el-container>
        </el-container>
    </el-container>
</template>

<script>
    import ProxyApi from './Pages/ProxyApi/index'

    export default {
        components: {
            ProxyApi
        },
        data() {
            return {
                isCollapse: false,
            };
        },
        mounted() {
            window.$app = this;
        }
    }
</script>

<style>
.app-header {
    line-height: 60px;
    font-size: 30px;
    color: white;
}
.el-menu-vertical-demo {
    height: calc(100% - 56px);
}
.app-main {
    height: 100%;
    overflow-y: hidden;
}
.app-el-menu {
    background: #f4f4f4;
}
.app-el-menu-li {
    border-bottom: 1px solid;
    height: 50px;
    line-height: 50px;
    font-size: 20px;
    display: block;
    width: 100%;
    text-align: center;
    cursor: pointer;
}

*::-webkit-scrollbar {
    /*滚动条整体样式*/
    width : 2px;  /*高宽分别对应横竖滚动条的尺寸*/
    height: 1px;
}
*::-webkit-scrollbar-thumb {
    /*滚动条里面小方块*/
    border-radius   : 10px;
    background-color: skyblue;
    background-image: -webkit-linear-gradient(
        45deg,
        rgba(255, 255, 255, 0.2) 25%,
        transparent 25%,
        transparent 50%,
        rgba(255, 255, 255, 0.2) 50%,
        rgba(255, 255, 255, 0.2) 75%,
        transparent 75%,
        transparent
    );
}
*::-webkit-scrollbar-track {
    /*滚动条里面轨道*/
    box-shadow   : inset 0 0 0px rgba(0, 0, 0, 0.2);
    background   : #ededed;
    border-radius: 10px;
}
</style>