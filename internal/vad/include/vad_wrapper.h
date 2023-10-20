#ifndef VAD_Wrapper_H
#define VAD_Wrapper_H

#ifdef __cplusplus
extern "C"
{
#endif

/**
* SDK使用者应实现VAD_Callback中的各函数，各函数指针的赋值不能为空！！！
* 注意：目前vad没有内建线程，VAD_Callback中的各函数是由VAD_Wrapper_FeedAudio函数调用的，所以不要在函数实现中进行耗时或者阻塞性操作！！！
*/
typedef struct
{
    /**
    * 语音开始
    * @param pUserData 应用层数据
    */
    void (*VoiceStart)(void* pUserData);
    /**
    * 语音结束
    * @param pUserData 应用层数据
    * @param info 关于本句语音的描述性信息（详见vad_end json协议），无需释放
    */
    void (*VoiceEnd)(void* pUserData, char* info);//语音结束
    /**
    * 语音数据
    * @param pUserData 应用层数据
    * @param data 语音数据
    * @param size 数据的大小
    * @param flags 0--正常语音 1--该语音段可能即将结束(可据此提前nlp处理)
    */
    void (*VoiceData)(void* pUserData, char *data, int size, int flags);
    /**
    * 非语音数据
    * @param pUserData 应用层数据
    * @param data 非语音数据
    * @param size 数据的大小
    */
    void (*NonVoiceData)(void* pUserData, char *data, int size);
}VAD_Callback;



/**
* 创建VAD_Wrapper实例
* @param f VAD_Callback
* @param pUserData 应用层数据
* @param mode 3---自研阵列，包括uac方式和line in方式（pad-pepper+rk3308+mic_array）
*             2---for ginger+微纳阵列
*             1---简单能量判据（针对pad-pepper+rk3308+mic_array，line in方式），目前已废弃，请改用mode3
*             0---其他
* @return VAD_Wrapper对象的句柄
*/
void *VAD_Wrapper_Create(VAD_Callback* f, void* pUserData, int mode);

/**
* 送入待处理的音频数据
* @param handle VAD_Wrapper对象的句柄
* @param buf 送入进行语音活性检测的数据，目前限定为16kHz 16bit mono PCM
* @param size 数据的大小，应不大于4096
* @param tips 仅用于自动化测试等目的，产品使用中应置为NULL
* @return &lt;0--失败， 0--当前为非语言帧，1--当前为语音帧
*/
int VAD_Wrapper_FeedAudio(void *handle, char* buf, int size, const char* tips);

/**
* 清空内部buffer,并重置状态到初始态
* @param handle VAD_Wrapper对象的句柄
*/
void VAD_Wrapper_Reset(void *handle);

/**
* 释放内部的所有资源
* @param handle VAD_Wrapper对象的句柄
*/
void VAD_Wrapper_Destory(void *handle);

/**
* 设置是否保存文件，未调用该接口的默认行为是不保存
* @param handle VAD_Wrapper对象的句柄
* @param save 0--不保存，1--保存
*/
void VAD_Wrapper_SetSaveFile(void *handle, char save);


#if 0
/**设置当前的声源位置(SSL,sound source location)
handle---VAD_Wrapper句柄
dimension 1--方位角 2--方位角+俯仰角 3--方位角(0~2pi)+俯仰角(0~pi)+径向距离(以米为单位)
size 目标个数
data dimension*size个float 顺序为：目标1的各维度，目标2的各维度...
*/
void VAD_Wrapper_SetCurrentSSL(void *handle, int dimension, int size, float* data);
/**设置期望的声源位置范围
handle---VAD_Wrapper句柄
dimension 1--方位角 2--方位角+俯仰角 3--方位角+俯仰角+径向距离
size 目标个数
data dimension*size*2 个float 顺序为：目标1的各维度范围（最小值，最大值），目标2的各维度范围（最小值，最大值）...
threshold 门限有效值0~1，在所观察的时间段内，当声源位置在该期望范围内的频率大于该门限时，认为该段语音是期望的有效语音，否则认为是无效语音
    &lt;0--代表使用内部默认的门限值.
*/
void VAD_Wrapper_SetDesiredSSL(void *handle, int dimension, int size, float* data, float threshold);
#endif

#ifdef __cplusplus
}
#endif

#endif //VAD_Wrapper_H

