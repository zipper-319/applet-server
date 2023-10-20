#ifndef VAD_WRAPPER_X86_EXTEND_H
#define VAD_WRAPPER_X86_EXTEND_H

#ifdef __cplusplus
extern "C"
{
#endif


void VAD_Wrapper_SetVadConfig(void* h, int enableVadAngleFilter, int dimension, int size, float* data, float threshold, int pauseMode, int enableEnergyFilter);



#ifdef __cplusplus
}
#endif

#endif //VAD_WRAPPER_X86_EXTEND_H

