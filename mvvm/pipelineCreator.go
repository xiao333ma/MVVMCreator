package mvvm

import "./utils"

type PipelineCreator struct{
	Creator
}

func (r PipelineCreator) CreateDotHFile() {
	content := `//
//  __prefix__Pipeline.h
//  __proj__
//
//  Created by __user__ on __time__.
//  Copyright © __year__ __user__. All rights reserved.
//

#import "IGCPipeline.h"

NS_ASSUME_NONNULL_BEGIN

@interface __prefix__Pipeline : IGCPipeline

@end

NS_ASSUME_NONNULL_END
`

	utils.WriteFile(utils.FileTypeOfPipeLine, true, content)

}

func (r PipelineCreator) CreateDotMFile() {
	content := `//
//  __prefix__Pipeline.m
//  __proj__
//
//  Created by __user__ on __time__.
//  Copyright © __year__ __user__. All rights reserved.
//

#import "__prefix__Pipeline.h"

@implementation __prefix__Pipeline

@end`

	utils.WriteFile(utils.FileTypeOfPipeLine, false, content)
}
